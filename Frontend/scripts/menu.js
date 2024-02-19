var menuButton = document.querySelector(".menu_button");
menuButton.addEventListener("click", (e) => {
    if (menuButton.classList.contains("menu_button__active")) {
        menuButton.classList.remove("menu_button__active")
    } else {
        menuButton.classList.add("menu_button__active")
    }
})