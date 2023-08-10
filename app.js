let tg = window.Telegram.WebApp;

tg.expand();

let btn1 = document.getElementById("btn1")

btn1.addListener("click", function(){
  
});

//Telegram.WebApp.onEvent("")
tg.sendData("qqq");

let p = document.createElement("p");

p.innerText = `${tg.initDataUnsafe.user.first_name}`
