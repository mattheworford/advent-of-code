import kotlin.math.abs

fun main() {
    fun isSafe1(levels: List<Int>) = when {
        levels.size <= 1 -> true
        else -> {
            val pairs = levels.zipWithNext()

            val allIncreasing = pairs.all { (a, b) -> b > a }
            val allDecreasing = pairs.all { (a, b) -> b < a }

            val validDifferences = pairs.all { (a, b) ->
                abs(b - a) in 1..3
            }

            (allIncreasing || allDecreasing) && validDifferences
        }
    }

    fun part1(input: List<String>): Int {
        return input.map { line ->
            isSafe1(line.split(" ").map { it.toInt() })
        }.count { it }
    }

    fun isSafe2(levels: List<Int>): Boolean {
        if (isSafe1(levels)) return true

        return levels.indices.any { i ->
            val withoutOne = levels.filterIndexed { index, _ -> index != i }
            isSafe1(withoutOne)
        }
    }

    fun part2(input: List<String>): Int {
        return input.map { line ->
            isSafe2(line.split(" ").map { it.toInt() })
        }.count { it }
    }

    val testInput = readInput("Day02_test")
    check(part1(testInput) == 2)
    check(part2(testInput) == 4)

    val input = readInput("Day02")
    part1(input).println()
    part2(input).println()
}