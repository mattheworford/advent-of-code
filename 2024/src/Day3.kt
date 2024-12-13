import kotlin.math.abs

fun main() {
    fun part1(input: List<String>): Int {
        val pattern = "mul\\((\\d{1,3}),(\\d{1,3})\\)".toRegex()

        return pattern.findAll(input.joinToString { it })
            .map { match ->
                // Extract the two numbers from the capture groups
                val (num1, num2) = match.destructured
                num1.toInt() * num2.toInt()
            }
            .sum()
    }

    fun part2(input: List<String>): Int {
        return -1
    }

    val testInput = readInput("Day03_test")
    check(part1(testInput) == 161)
    check(part2(testInput) == -1)

    val input = readInput("Day03")
    part1(input).println()
    part2(input).println()
}