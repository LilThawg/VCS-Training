const passwordInpDOM = document.querySelector("#password");
const emailInpDOM = document.querySelector("#email");
const buttonDOM = document.querySelector("#login-btn");
const notifyDOM = document.querySelector(".notification");
const login = async () => {
  const password = passwordInpDOM.value;
  const email = emailInpDOM.value;
  const postReg = { Email: email, Password: password };
  console.log(postReg);
  const res = await fetch("http://localhost:4000/login", {
    method: "POST", // *GET, POST, PUT, DELETE, etc.
    mode: "cors", // no-cors, *cors, same-origin
    credentials: "same-origin", // include, *same-origin, omit
    headers: {
      "Content-Type": "application/json",
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: JSON.stringify(postReg),
  });
  const data = await res.json();
  console.log(data);
  if (data.token) {
    console.log(`${data.role} login successfully`);
    localStorage.setItem("user", JSON.stringify(data));
    notifyDOM.innerHTML = "Login sucessfully";
    notifyDOM.classList.remove("error");
    notifyDOM.classList.add("show", "success");

    // redirect to home
    setTimeout(() => {
      location.assign("http://127.0.0.1:5500/Auth-jwt/FrontEnd/Home/Home.html");
    }, 2000);
  } else {
    console.log(`login failed`);
    notifyDOM.innerHTML = "Login failed check your credentials";
    notifyDOM.classList.remove("success");
    if (!notifyDOM.classList.contains("show")) {
      notifyDOM.classList.add("show");
    }
    notifyDOM.classList.add("error");
  }
};

buttonDOM.addEventListener("click", login);
