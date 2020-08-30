const player = [0, 1];
const winConditions = ["123", "456", "789", "147", "258", "369", "159", "357"];
let turn = 0;
const color = {
    player1: ["red", "blue", "green"],
    player2: ["red", "blue", "green"],
};
let gameState = false;
let state = ["", "", "", "", "", "", "", "", ""];
const DOMs = {
    box: document.querySelector(".box"),
    headerText: document.querySelector(".header__text"),
    restartBtn: document.querySelector(".restart__btn"),
};
let count = 0;
const init = () => {
    count = 0;
    turn = 0;
    gameState = false;
    state = ["", "", "", "", "", "", "", "", ""];
    DOMs.headerText.textContent = "O Player's Turn";
    DOMs.headerText.classList.remove(color.player1[0]);
    DOMs.headerText.classList.add(color.player1[0]);
    DOMs.headerText.classList.remove(color.player2[1]);
    const squareBoxes = Array.from(document.querySelectorAll(".box__item--square"));
    squareBoxes.forEach((box, index) => {
        box.classList.remove(color.player1[0]);
        box.classList.remove(color.player2[1]);
        box.textContent = index + 1;
    });
};
const checkResult = (state, type) => {
    let answer = "";
    for (let i = 0; i < 8; i++) {
        switch (i) {
            case 0:
                answer = `${state[0]}${state[1]}${state[2]}`;
                break;
            case 1:
                answer = `${state[3]}${state[4]}${state[5]}`;
                break;
            case 2:
                answer = `${state[6]}${state[7]}${state[8]}`;
                break;
            case 3:
                answer = `${state[0]}${state[3]}${state[6]}`;
                break;
            case 4:
                answer = `${state[1]}${state[4]}${state[7]}`;
                break;
            case 5:
                answer = `${state[2]}${state[5]}${state[8]}`;
                break;
            case 6:
                answer = `${state[0]}${state[4]}${state[8]}`;
                break;
            case 7:
                answer = `${state[2]}${state[4]}${state[6]}`;
                break;
        }
        console.log(answer);
        if (answer === "OOO") {
            return "O";
        } else if (answer === "XXX") {
            return "X";
        }
    }
    count += 1;
    console.log(count);
    if (count === 9) {
        return "DRAW";
    } else {
        return "GO";
    }
};

DOMs.box.addEventListener("click", (e) => {
    if (!gameState) {
        const element = e.target;
        const elementId = e.target.dataset.id;
        if (elementId && elementId !== "0" && elementId !== "10") {
            if (turn === 0) {
                if (element.textContent !== "O" && element.textContent !== "X") {
                    state[+elementId - 1] = "O";

                    element.textContent = "O";
                    element.classList.add(color.player1[0]);
                    let result = checkResult(state, "O");
                    if (result === "O") {
                        DOMs.headerText.textContent = "O Player wins!!!";

                        gameState = true;
                    } else if (result === "DRAW") {
                        DOMs.headerText.textContent = "DRAW, Play Again!!!";
                        DOMs.headerText.classList.remove(color.player1[0]);
                        DOMs.headerText.classList.remove(color.player2[1]);
                        gameState = true;
                    } else {
                        DOMs.headerText.textContent = "X Player's Turn";
                        DOMs.headerText.classList.toggle(color.player1[0]);
                        DOMs.headerText.classList.toggle(color.player2[1]);

                        turn = 1;
                    }
                }
            } else {
                if (element.textContent !== "O" && element.textContent !== "X") {
                    state[+elementId - 1] = "X";
                    element.textContent = "X";

                    element.classList.add(color.player2[1]);
                    let result = checkResult(state, "X");
                    if (result === "X") {
                        DOMs.headerText.textContent = "X Player wins!!!";

                        gameState = true;
                        turn = 0;
                    } else if (result === "DRAW") {
                        DOMs.headerText.textContent = "DRAW, Play Again!!!";
                        DOMs.headerText.classList.remove(color.player1[0]);
                        DOMs.headerText.classList.remove(color.player2[1]);

                        gameState = true;
                    } else {
                        DOMs.headerText.textContent = "O Player's Turn";
                        DOMs.headerText.classList.toggle(color.player1[0]);
                        DOMs.headerText.classList.toggle(color.player2[1]);
                        turn = 0;
                    }
                }
            }
        }
    }
});

DOMs.restartBtn.addEventListener("click", init);
