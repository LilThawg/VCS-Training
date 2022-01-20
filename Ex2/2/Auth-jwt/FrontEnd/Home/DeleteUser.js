const emailInpDOM = document.querySelector('#email')
const buttonDOM = document.querySelector('#deleteuser-btn')

user = JSON.parse(localStorage.getItem("user"));

const deleteuser = async () => {
    const email = emailInpDOM.value
    const postReg = {'Email': email}
    console.log(postReg)
    const res = await fetch('http://localhost:4000/deleteuser',{
        method: 'DELETE', // *GET, POST, PUT, DELETE, etc.
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

buttonDOM.addEventListener('click', deleteuser);