import java.util.*
import kotlin.math.abs

fun main() {
    fun part1(input: List<String>): Int {
        val (left, right) = input.map { locations ->
            val (left, right) = locations.split("   ").map { it.toInt() }
            left to right
        }.unzip()
        val leftLocations = PriorityQueue<Int>().apply { addAll(left) }
        val rightLocations = PriorityQueue<Int>().apply { addAll(right) }
        return generateSequence {
            if (leftLocations.isNotEmpty())
                abs(rightLocations.poll() - leftLocations.poll())
            else null
        }.sum()
    }

    fun part2(input: List<String>): Int {
        val (leftLocations, rightLocations) = input.map { locations ->
            val (left, right) = locations.split("   ").map { it.toInt() }
            left to right
        }.unzip()
        val counter = rightLocations.groupingBy { it }.eachCount()
        return leftLocations.sumOf { counter[it]?.times(it) ?: 0 }
    }

    val testInput = readInput("Day01_test")
    check(part1(testInput) == 11)
    check(part2(testInput) == 31)

    val input = readInput("Day01")
    part1(input).println()
    part2(input).println()
}