function startLoadTest() {
  fetch("/api/v1/start", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });
}

document
  .getElementById("startLoadTestBtn")
  .addEventListener("click", startLoadTest);
