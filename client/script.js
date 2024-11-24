// connect to the websocket
const socket = new WebSocket("http://localhost:8080/connect")
socket.onopen = () => {
  console.log("Connected!")
}

const canvas = document.getElementById("game-canvas");
const ctx = canvas.getContext('2d');
ctx.imageSmoothingEnabled = false;
const playerImg = document.createElement('img')
playerImg.src = 'santa_mk1.png'

socket.onmessage = (event) => {
  ctx.clearRect(0, 0, canvas.clientWidth, canvas.clientHeight)
  const entities = JSON.parse(event.data)
  for(entity of entities) {
    ctx.drawImage(playerImg, entity.X, -entity.Y, 64, 64)
  }

}

let controlls = {x: 0, y: 0};
window.addEventListener("keyup", () => {
  socket.send(JSON.stringify(controlls))
}) 

window.addEventListener("keydown", (e) => {

  switch(e.key) {
    case 'w': {
      controlls.y = 1;
      break;
    }
    case 'a': {
      controlls.x = -1;
      break;
    }
    case 's': {
      controlls.y = -1;
      break;
    }
    case 'd': {
      controlls.x = 1;
      break;
    }
    default: {
      // nothing
    }
  }
  console.log(controlls)
  socket.send(JSON.stringify(controlls))
  controlls = {x: 0, y: 0};
})