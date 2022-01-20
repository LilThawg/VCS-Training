const getAllDom = document.getElementById("get-all-btn");
const addDom = document.getElementById("add-btn");
const deleteDom = document.getElementById("delete-btn");
const logoutDom = document.getElementById("logout-btn");

const showDataDOM = document.getElementById("show-data-container");
user = JSON.parse(localStorage.getItem("user"));
const welcomeDOM = document.getElementById("welcome");
if (user?.email && user?.token) {
  welcomeDOM.innerText = `Welcome ${user.email}`;
  if (user.role === "admin") {
    getAllDom.addEventListener("click", async (e) => {
      e.preventDefault();
      const res = await fetch("http://localhost:4000/getalluser", {
        method: "GET", // *GET, POST, PUT, DELETE, etc.
        mode: "cors", // no-cors, *cors, same-origin
        credentials: "same-origin", // include, *same-origin, omit
        headers: {
          "Content-Type": "application/json",
          token: user.token,
        },
      });
      const data = await res.json();
      console.log(data);
      showDataDOM.innerText = data
        .map(
          (user) => `
            id: ${user.ID}, 
            created_at: ${user.CreatedAt},
            updated_at:${user.UpdatedAt},
            name:${user.name},
            email:${user.email},
            role:${user.role}
      `
        )
        .join("\n");
    });
  }

  // logout
  logoutDom.addEventListener("click", (e) => {
    e.preventDefault();
    localStorage.removeItem("user");
    location.assign("http://127.0.0.1:5500/Auth-jwt/FrontEnd/Home/index.html");
  });
} else {
  welcomeDOM.innerText = "Not authorized. Please login";
  location.assign("http://127.0.0.1:5500/Auth-jwt/FrontEnd/Home/index.html");
}
