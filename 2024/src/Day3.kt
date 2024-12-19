import kotlin.math.abs

fun main() {
    fun part1(input: List<String>): Int {
        val pattern = "mul\\((\\d{1,3}),(\\d{1,3})\\)".toRegex()

        return pattern.findAll(input.joinToString { it })
            .map { match ->
                val (num1, num2) = match.destructured
                num1.toInt() * num2.toInt()
            }
            .sum()
    }

    fun part2(input: List<String>): Int {
        val mulPattern = "mul\\((\\d{1,3}),(\\d{1,3})\\)".toRegex()
        val doPattern = "do\\(\\)".toRegex()
        val dontPattern = "don't\\(\\)".toRegex()

        var enabled = true
        var sum = 0

        val allMatches = (mulPattern.findAll(input.joinToString { it }) +
                doPattern.findAll(input.joinToString { it }) +
                dontPattern.findAll(input.joinToString { it }))
            .sortedBy { it.range.first }

        for (match in allMatches) {
            when (match.value) {
                "do()" -> enabled = true
                "don't()" -> enabled = false
                else -> {
                    if (enabled) {
                        val (num1, num2) = match.destructured
                        sum += num1.toInt() * num2.toInt()
                    }
                }
            }
        }
        return sum
    }

    val testInput1 = readInput("Day03_test_part1")
    check(part1(testInput1) == 161)
    val testInput2 = readInput("Day03_test_part2")
    check(part2(testInput2) == 48)

    val input = readInput("Day03")
    part1(input).println()
    part2(input).println()
}