//handle when the number is pressed. It renders the number into #input element
function number(number) {
    //beginanswer
    let input = document.getElementById('input');
    input.value += number;
    //endanswer
}

//handle when the Del button, AC button and other operators (+, /, -, x) are pressed. It renders the operator into #input element
function operator(operator) {
    let input = document.getElementById('input');
    if (operator == 'Del') {
        //beginanswer
        input.value = input.value.slice(0, -1);
        //endanswer
    } else if (operator == "AC") {
        //beginanswer
        input.value = '';
        //endanswer
    } else {
        //beginanswer
        input.value += operator;
        //endanswer
    }
}

//handle when = button is pressed. It renders the result into #input element
function calculate() {
    //beginanswer
    let input = document.getElementById('input');
    input.value = eval(input.value);
    //endanswer
}