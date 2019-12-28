var actualNumber = Math.floor(Math.random() * 10 + 1)

var numberOfguesses = 1;

document.getElementById("submitguess").onclick = function() {
    var guess = document.getElementById("guessField").value;
    
    if (guess == actualNumber) {
        alert("You guessed the right number. It took you " + numberOfguesses + " guesses.");
    }
    else if (guess > actualNumber) {
        numberOfguesses++;
        alert("Incorrect guess. The actual number is lower.");
    }
    else if (guess < actualNumber) {
        numberOfguesses++;
        alert("Incorrect guess. The actual number is higher.");
    }
}