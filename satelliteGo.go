/*
 * author: Andrea Daza
 * email: andreacdazar1@gmail.com
 * topic: Operación fuego de Quasar
 * date: 07/06/2021
 */
package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// posiciones determinadas de los 3 satélites
var Positions = [][]float64{{-500, -200}, {100, -100}, {500, 100}}

type Satellite struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Response struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

// triangula la posición de la nave
func GetLocation(distances ...float64) (x, y float64, err bool) {
	// control de cantidad de distancias vs posiciones
	if len(distances) == len(Positions) {
		allPoints := make([][][]float64, 0)
		for i, distance := range distances {
			if i == 0 {
				continue
			}
			// obtiene los puntos de intersección del satélite 1 con los demás
			allPoints = append(allPoints, GetInterPoints(distances[0], distance, Positions[0], Positions[i]))
		}
		intersection := Compare(allPoints)
		if intersection != nil {
			return intersection[0], intersection[1], false
		} else {
			fmt.Println("Lo siento, los puntos de intersección no coinciden. ")
		}
	} else {
		fmt.Println("Se deben ingresar los datos de los satélites configurados:", len(Positions))
	}
	return 0, 0, true
}

// obtiene los puntos de intersección repetidos
func Compare(allPoints [][][]float64) []float64 {
	for i, points := range allPoints {
		for j := i + 1; j < len(allPoints); j++ {
			if Equal(points[0], allPoints[j]) {
				return points[0]
			} else if Equal(points[1], allPoints[j]) {
				return points[1]
			}
		}
	}
	return nil
}

// compara arrays de puntos de intersección
func Equal(point []float64, points [][]float64) bool {
	return reflect.DeepEqual(point, points[0]) || reflect.DeepEqual(point, points[1])
}

// Calcula los puntos de intersección de dos esferas
func GetInterPoints(d0 float64, d1 float64, p0 []float64, p1 []float64) [][]float64 {
	// ecuacion a
	xa := p0[0] * -2.0
	ya := p0[1] * -2.0
	ca := GetC(d0, p0[0], p0[1])
	// ecuacion b
	xb := p1[0] * -2.0
	yb := p1[1] * -2.0
	cb := GetC(d1, p1[0], p1[1])
	// ecuacion solucion
	xs := xa - xb
	ys := (ya - yb) * -1
	cs := ca - cb

	xdc := cs / xs
	xdy := ys / xs

	// ecuacion final
	a := math.Pow(xdy, 2) + 1             // x^2
	b := 2*(xdy)*xdc + (xb * xdy) + yb    // 2x
	c := (xb*xdc + math.Pow(xdc, 2)) - cb // 2
	// Y
	y1, y2 := Quadratic(a, b, c)
	// X
	x1 := (ys*y1 + cs) / xs
	x2 := (ys*y2 + cs) / xs

	return [][]float64{{Round(x1), Round(y1)}, {Round(x2), Round(y2)}}
}

// Redondea la posición a dos cifras decimales
func Round(n float64) float64 {
	return math.Round(n*100) / 100
}

// obtiene la constante
func GetC(distance, x, y float64) float64 {
	return math.Pow(distance, 2) - (math.Pow(x, 2) + math.Pow(y, 2))
}

// Ecuación cuadrática
func Quadratic(a, b, c float64) (y1, y2 float64) {
	b24ac := math.Pow(b, 2) - 4*(a*c)
	y1 = (-b + math.Sqrt(b24ac)) / (2 * a)
	y2 = (-b - (math.Sqrt(b24ac))) / (2 * a)
	return y1, y2
}

// Obtiene mensaje resultante
func GetMessage(messages ...[]string) (mssg string, err bool) {
	var msg []string
	for i := 0; i < len(messages[0]); i++ {
		for j := 0; j < len(messages); j++ {
			// control longitud mensajes
			if len(messages[0]) != len(messages[j]) {
				return "", true
			}
			// control de mensajes vacios y repetidos
			if messages[j][i] != "" && (len(msg) == 0 || msg[len(msg)-1] != messages[j][i]) {
				msg = append(msg, messages[j][i])
			}
		}
	}
	// control mensaje vacio
	if len(msg) > 0 {
		return strings.Join(msg, " "), false
	} else {
		return "", true
	}
}
