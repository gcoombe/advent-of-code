package main

import (
	"runtime"
	"path"
	"fmt"
	"advent-of-code/2017/goUtils"
	"bufio"
	"math"
	"regexp"
	"strconv"
	"strings"
)


func main() {

	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "input.txt")

	fmt.Println("----------------------------------------")
	fmt.Println("DAY 20")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %s", runPartOne(filepath))
	fmt.Printf("\nPart 2 res: %d", runPartTwo(filepath))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 19")
	fmt.Println("----------------------------------------")
}

type Particle struct {
	id int
	p []int
	v []int
	a []int
	dead bool
}

func (particle *Particle) tick() {
	for i := 0; i < 3; i++ {
		particle.v[i] += particle.a[i]
	}
	for i := 0; i < 3; i++ {
		particle.p[i] += particle.v[i]
	}
}

func (particle *Particle) distance() float64 {
	return math.Abs(float64(particle.p[0])) + math.Abs(float64(particle.p[1])) + math.Abs(float64(particle.p[2]))
}

func (particle *Particle) collide(otherParticle *Particle) bool {
	return particle.p[0] == otherParticle.p[0] && particle.p[1] == otherParticle.p[1] && particle.p[2] == otherParticle.p[2]
}

func parseTuple(input string) []int {
	tuple := make([]int, 3)
	for i, str := range strings.Split(input, ",") {
		tuple[i], _ = strconv.Atoi(str)
	}
	return  tuple
}

func buildParticles(filePath string) []*Particle {
	file := goUtils.GetFile(filePath)

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	regex := `p=\<([\d, -]*)\>, v=\<([\d, -]*)\>, a=\<([\d, -]*)\>`
	particles := make([]*Particle, 0)

	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		matchingRegex := regexp.MustCompile(regex)
		matches := matchingRegex.FindStringSubmatch(line)

		particles = append(particles, &Particle{i,parseTuple(matches[1]), parseTuple(matches[2]), parseTuple(matches[3]), false})
		i++
	}

	return particles
}

func runPartOne(filePath string) int {
	particles := buildParticles(filePath)

	closest := particles[0]
	for i := 0; i < 10000; i++ {
		for _, particle := range particles {
			particle.tick()
			if particle.distance() < closest.distance() {
				closest = particle
			}

		}
	}
	return closest.id
}

func runPartTwo(filePath string) int {
	particles := buildParticles(filePath)

	for i := 0; i < 10000; i++ {
		for _, particle := range particles {
			particle.tick()
		}

		for i, particle := range particles {
			particlesToKill := make([]*Particle, 0)
			for _, otherParticle := range particles[i+1:] {
				if particle.dead == false && otherParticle.dead == false && particle.collide(otherParticle) {
					particlesToKill = append(particlesToKill, particle, otherParticle)
				}
			}
			for _, particleIToKill := range particlesToKill {
				particleIToKill.dead = true
			}

		}
	}
	particlesLeft := 0
	for _, particle := range particles {
		if particle.dead == false {
			particlesLeft++
		}
	}
	return  particlesLeft
}