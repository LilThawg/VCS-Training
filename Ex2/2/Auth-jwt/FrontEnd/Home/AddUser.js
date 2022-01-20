const usernameInpDOM = document.querySelector('#username')
const passwordInpDOM = document.querySelector('#password')
const emailInpDOM = document.querySelector('#email')
const roleInpDOM = document.querySelector('#role')
const buttonDOM = document.querySelector('#adduser-btn')

user = JSON.parse(localStorage.getItem("user"));

const adduser = async () => {
    const username = usernameInpDOM.value
    const password = passwordInpDOM.value
    const email = emailInpDOM.value
    const role = roleInpDOM.value
    const postReg = {'Name': username, 'Password': password, 'Email': email, 'Role': role}
    console.log(postReg)
    const res = await fetch('http://localhost:4000/adduser',{
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
          'Content-Type': 'application/json',
          // 'Content-Type': 'application/x-www-form-urlencoded',
          token: user.token,
        },
        body: JSON.stringify(postReg)
     }
        );
    const data = await res.json()
    console.log(data)
}

buttonDOM.addEventListener('click', adduser);