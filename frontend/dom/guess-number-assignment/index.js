/*
Uncomment variable dibawah ini untuk mulai mengerjakan. dilarang mengganti nama variable yang dibuat.
*/

// let newGameBtn = tombol untuk melakukan restart game
// let randomNumber = generate random number dari 1 sampai 10
// let message = element untuk menampilkan pesan
// let displayScore = element untuk menampilkan score
// let input = element untuk memasukan data
// let checkBtn = tombol untuk melakuan pengecekan angka pada input
// let score = nilai yang akan ditampilakan

<<<<<<< HEAD
//beginanswer
let newGameBtn = document.getElementsByClassName('new-game')[0]
let randomNumber = Math.floor((Math.random() * 10) + 1)
let message = document.getElementsByClassName('message')[0]
let displayScore = document.getElementsByClassName('score')[0]
let input = document.getElementsByClassName('input')[0]
let checkBtn = document.getElementsByClassName('check')[0]
let score = 10

newGameBtn.addEventListener('click', resetGame)
checkBtn.addEventListener('click', checkNumber)
//endanswer


function displayMessage(msg) {
  //beginanswer
  message.innerHTML = msg
  //endanswer
}

function resetGame() {
  //beginanswer
  score = 10
  displayScore.innerHTML = 10
  input.value = ''
  displayMessage('Lets start guessing...')
  document.getElementsByClassName('number')[0].innerHTML = '?'
  randomNumber = Math.floor((Math.random() * 10) + 1)
  console.log(randomNumber);
  //endanswer
=======
// TODO: answer here


function displayMessage(msg) {
  // TODO: answer here
}

function resetGame() {
  // TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
  //dilarang menghapus code dibawah ini!
  document.getElementById('hidden-number').innerHTML = randomNumber
}

function checkNumber() {
<<<<<<< HEAD
  //beginanswer
  if (input.value < 1 || input.value > 10) {
    displayMessage('Guess any number between 1 and 10')
  } else if (input.value == randomNumber) {
    displayMessage('Yeay! you guessed it!')
    document.getElementsByClassName('number')[0].innerHTML = randomNumber
  } else if (input.value != randomNumber) {
    if (score > 1) {
      if (input.value < randomNumber) {
        displayMessage('Too small, buddy!')
      } else if (input.value > randomNumber) {
        displayMessage("Oww... that's too big!")
      }
      score--
      displayScore.innerHTML = score
    } else {
      displayMessage('Oops, you lost the game')
      displayScore.innerHTML = 0
    }
  }
  //endanswer
=======
  // TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

//dilarang menghapus code dibawah ini!
document.getElementById('hidden-number').innerHTML = randomNumber