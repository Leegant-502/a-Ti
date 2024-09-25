console.log(document.getElementById("reg"));
function forget() { alert("不会做，hh（六花！六花！）"); }


function login(event) {
    event.preventDefault();
    const password = document.getElementsByClassName("password")[0].value;
    const name = document.getElementsByClassName("name")[0].value;
    const factory = [{ name: "1229761829", password: "123456" },
    { name: "neko", password: "114514" },
    { name: "taffy", password: "1919810" },
    { name: "qiye", password: "0426" },
    ]




    const user = factory.find(user => user.name === name && user.password === password);

    if (user) {
        alert("登录成功,即将进入提瓦特大陆");
        window.location.href = 'https://ys.mihoyo.com/';
    }
    else {
        alert("用户名或密码错误！");
        ;
    }
}

reg.addEventListener("click", function regg() {

    if
        (registerform.style.display === "none") { registerform.style.display = "flex"; loginform.style.display = "none"; }
    else { registerform.style.display = "none"; }
})


document.getElementById('button1').addEventListener("click",
    function reggg() {
        const name = document.getElementsByClassName("name1")[0].value;
        const password1 = document.getElementsByClassName("password1")[0].value;
        const password2 = document.getElementsByClassName("password2")[0].value;
        if
            (name == "" || password1 == "" || password2 == "") {
            alert("为什么不填？？？，虽然填了也没有用");
            return;
        }
        else if (password1 !== password2) {
            alert("看好了再输入，两次密码不同");
            return;
        }
        else {
            alert("注册成功，但是依然登陆不了，hhhhhhhhhh");
        }
    })

function checkpassword() {
    const password = document.getElementById("password");
    if (password.type === "password"){
        password.type = "text";
    }
    else {
        password.type = "password";
    }
}
const checkPassword=document.getElementById("checkPassword");
checkPassword.addEventListener("click", checkpassword);