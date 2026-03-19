document.addEventListener("DOMContentLoaded", () => main());
const main = () => {
  const jsonOutput = document.querySelector("#jsonOutput");
  const jsonInput = document.querySelector("#jsonInput");
  const formatJson = (event) => {
    console.log("here");
    console.log(event.target.value);
    res = FormatJson(event.target.value);
    jsonOutput.value = res;
  };

  jsonInput.addEventListener("input", formatJson);
  function escapeHTML(str) {
    return str
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;");
  }

  function highlightJSON(json) {
    json = escapeHTML(json);

    return json
      .replace(/("(.*?)")(?=:)/g, '<span class="key">$1</span>')
      .replace(/:\s*("(.*?)")/g, ': <span class="string">$1</span>')
      .replace(/\b(true|false|null)\b/g, '<span class="bool">$1</span>')
      .replace(/\b\d+(\.\d+)?\b/g, '<span class="number">$&</span>');
  }

  function handleInput(value) {
    const highlight = document.getElementById("highlight");

    highlight.innerHTML = highlightJSON(value);

    // your existing WASM formatter
    if (typeof formatJson === "function") {
      formatJson(value);
    }
  }
};
