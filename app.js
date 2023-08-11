let tg = window.Telegram.WebApp;

tg.expand();

tg.MainButton.textColor = '#FFFFFF';
tg.MainButton.color = '#2cab37';

let user = tg.initDataUnsafe.user;

let btn1 = document.getElementById("btn1")

//btn1.addListener("click", function(){
//
//});

//Telegram.WebApp.onEvent("")
tg.sendData("qqq");

let usercard = document.getElementById("usercard");


if(user === undefined){
   var e = document.getElementsByTagName('html')[0];
   e.removeChild(document.body);
}

let p = document.createElement("p");
p.innerText = user;//`${tg.initDataUnsafe.user.first_name}
//${tg.initDataUnsafe.user.last_name}`;

usercard.appendChild(p);
