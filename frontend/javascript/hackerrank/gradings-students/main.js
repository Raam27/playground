// Grading Students
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

// Full Problem: https://www.hackerrank.com/challenges/grading/problem

function gradingStudents(grades) {
  // Write your code here
<<<<<<< HEAD
  //beginanswer
  var newGrades = [];
  for (var i = 0; i < grades.length; i++) {
    if (grades[i] >= 38) {
      if (grades[i] % 5 >= 3) {
        newGrades.push(grades[i] - (grades[i] % 5) + 5);
      } else {
        newGrades.push(grades[i]);
      }
    } else {
      newGrades.push(grades[i]);
    }
  }
  return newGrades
  //endanswer
=======
  // TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

function main() {
  //     var n = parseInt(readLine());
  //     var grades = [];
  //     for(var grades_i = 0; grades_i < n; grades_i++){
  //        grades[grades_i] = parseInt(readLine());
  //     }
  var grades = [73, 67, 38, 33]; // override input
  var result = gradingStudents(grades);
  console.log(result.join("\n"));
}

main(); // execute program

module.exports = gradingStudents