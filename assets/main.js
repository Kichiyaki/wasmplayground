const addEventListeners = () => {
  document.querySelector("form").addEventListener("submit", (e) => {
    e.preventDefault();

    const a = parseFloat(e.target.a.value);
    const b = parseFloat(e.target.b.value);

    document.querySelector("#result").innerHTML = add(a, b);
  });
};

const run = () => {
  addEventListeners();
};

const go = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
  (res) => {
    go.run(res.instance);

    run();
  }
);
