package algorithm

import (
	"log"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func FindClosestPair(pointsX, pointsY []*Point) float64 {
	log.Printf("Finding closest pair in: %v", getPoint(pointsX))
	if len(pointsX) <= 3 {
		minimumDistance := findMinimumDistance(pointsX)
		log.Printf("Minimum distance is: %v", minimumDistance)
		return minimumDistance
	}

	distance1 := FindClosestPair(pointsX[:len(pointsX)/2], pointsY)
	distance2 := FindClosestPair(pointsX[len(pointsX)/2:], pointsY)
	minimumDistance := math.Min(distance1, distance2)

	var strip []*Point
	xOfMiddleElement := pointsX[len(pointsX)/2].X
	for _, point := range pointsY {
		if point.X >= xOfMiddleElement-minimumDistance && point.X <= xOfMiddleElement+minimumDistance {
			strip = append(strip, point)
		}
	}
	log.Printf("The strip is: %v of length: %v", getPoint(strip), len(strip))
	minimumDistance = findSplitPair(strip, minimumDistance)

	log.Printf("Minimum distance is: %v", minimumDistance)
	return minimumDistance
}

func findSplitPair(points []*Point, currentMin float64) float64 {
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points) && math.Abs(points[j].Y-points[i].Y) < currentMin; j++ {
			currentMin = math.Min(findEuclideanDistance(points[i], points[j]), currentMin)
		}
	}
	return currentMin
}

func findMinimumDistance(points []*Point) float64 {
	var minimumDistance float64
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			distance := findEuclideanDistance(points[i], points[j])
			if minimumDistance == 0 {
				minimumDistance = distance
			} else {
				minimumDistance = math.Min(minimumDistance, distance)
			}
		}
	}
	return minimumDistance
}

func findEuclideanDistance(p1, p2 *Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	distance := math.Sqrt((dx * dx) + (dy * dy))
	log.Printf("Distance between points %v and %v is %v", p1, p2, distance)
	return distance
}

func getPoint(pointerArray []*Point) []Point {
	var points []Point
	for _, point := range pointerArray {
		points = append(points, *point)
	}
	return points
}
