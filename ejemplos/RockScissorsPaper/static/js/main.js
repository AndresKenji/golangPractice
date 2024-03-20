console.log("Cargando sabrosura")

let imgArray = [
    "static/img/hand_rock.png",
    "static/img/hand_paper.png",
    "static/img/hand_scissors.png"
];


function choose(x) {
    fetch("/play?c="+x)
    .then(r => r.json())
    .then(data => {
        console.log(data)
        if (x == 0) {
            document.getElementById("player_choice").innerHTML = "El jugador eligio PIEDRA";
        }else if (x ==1){
            document.getElementById("player_choice").innerHTML = "El jugador eligio PAPEL";
        } else {
            document.getElementById("player_choice").innerHTML = "El jugador eligio TIJERA";
        }

        document.getElementById("player_score").innerHTML = data.player_score;
        document.getElementById("computer_score").innerHTML = data.computer_score;
        document.getElementById("computer_choice").innerHTML = data.computer_choice;
        document.getElementById("round_result").innerHTML = data.round_result;
        document.getElementById("round_message").innerHTML = "Computadora: "+ data.message;

        var imgElement = document.getElementById("img_computer");
        imgElement.src = imgArray[data.computer_choiceInt]
    })
    .catch(err => {
        console.log(err);
    })
}