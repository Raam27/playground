// A Very Big Sum
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

// Full Problem: https://www.hackerrank.com/challenges/a-very-big-sum/problem

function aVeryBigSum(ar) {
  // Complete this function
<<<<<<< HEAD
  //beginanswer
  let sum = 0;
    for (var i = 0; i < ar.length; i++) {
        sum += ar[i];
    }
    return sum;
  //endanswer
=======
  // TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

function main() {
  var ar = [1000000001, 1000000002, 1000000003, 1000000004, 1000000005]; // override input
  var result = aVeryBigSum(ar);
  console.log(result);
}

main(); // execute program

module.exports = aVeryBigSum