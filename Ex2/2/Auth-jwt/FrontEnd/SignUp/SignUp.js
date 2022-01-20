const usernameInpDOM = document.querySelector('#username')
const passwordInpDOM = document.querySelector('#password')
const emailInpDOM = document.querySelector('#email')
const buttonDOM = document.querySelector('#signup-btn')
const signup = async () => {
    const username = usernameInpDOM.value
    const password = passwordInpDOM.value
    const email = emailInpDOM.value
    const postReg = {'Name': username, 'Password': password, 'Email': email}
    console.log(postReg)
    const res = await fetch('http://localhost:4000/signup',{
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
          'Content-Type': 'application/json'
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: JSON.stringify(postReg)
     }
        );
    const data = await res.json()
    console.log(data)
}

buttonDOM.addEventListener('click', signup);
