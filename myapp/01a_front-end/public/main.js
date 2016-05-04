var el = document.querySelector
  ("#enter-area");

el.addEventListener("input",
  function(e) {
    console.log(e.target.value);
  })
