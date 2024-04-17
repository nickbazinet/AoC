package main

import (
    "math"
)

func solveEquations(additionResult, multiplicationResult float64) (float64, float64) {
    // Coefficients for the quadratic equation derived from the given system of equations
    a := 1.0
    b := math.Abs(additionResult) * -1
    c := multiplicationResult

    // Calculate the discriminant
    discriminant := math.Sqrt(b*b - 4*a*c)

    // Calculate the solutions using the quadratic formula
    s1 := (-b + discriminant) / (2 * a)
    s2 := (-b - discriminant) / (2 * a)

    return s1, s2
}
