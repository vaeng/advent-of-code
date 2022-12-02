defmodule AdventOfCode.Day02 do
  def part1(args) do
    convertToLine(args)
    |> Enum.map(&getScore(&1))
    |> Enum.sum()
  end

  def part2(args) do
    convertToLine(args)
    |> Enum.map(&getScore2(&1))
    |> Enum.sum()
  end

  defp convertToLine(str),
    do:
      str
      |> String.trim()
      |> String.split("\n")

  # A Rock, B Paper, C Scissors
  # X Rock, Y Paper, Z Scissors
  # 1 Rock, 2 Paper, 3 Scissors
  # 6 Win, 3 Draw, 0 Loss

  # rock-paper -> Win
  defp getScore("A Y"), do: 2 + 6
  # paper-rock -> Loss
  defp getScore("B X"), do: 1 + 0
  # Scissors-Scissors -> Draw
  defp getScore("C Z"), do: 3 + 3

  # paper paper -> draw
  defp getScore("B Y"), do: 2 + 3
  # scis rock -> win
  defp getScore("C X"), do: 1 + 6
  # rock scis -> loss
  defp getScore("A Z"), do: 3 + 0

  # scis -> paper -> loss
  defp getScore("C Y"), do: 2 + 0
  # rock - rock -> draw
  defp getScore("A X"), do: 1 + 3
  # paper - scis -> win
  defp getScore("B Z"), do: 3 + 6

  defp getScore(_else), do: "Error"

  # A Rock, B Paper, C Scissors
  # X Lose, Y Draw, Z Win
  # 1 Rock, 2 Paper, 3 Scissors
  # 6 Win, 3 Draw, 0 Loss

  # rock - draw -> rpck
  defp getScore2("A Y"), do: 3 + 1
  # paper - loss -> rock
  defp getScore2("B X"), do: 0 + 1
  # scis - win -> rock
  defp getScore2("C Z"), do: 6 + 1

  # paper - draw -> paper
  defp getScore2("B Y"), do: 3 + 2
  # scis - loss
  defp getScore2("C X"), do: 0 + 2
  # rock - win
  defp getScore2("A Z"), do: 6 + 2

  # scis - draw
  defp getScore2("C Y"), do: 3 + 3
  # rock - loss
  defp getScore2("A X"), do: 0 + 3
  # paper - win
  defp getScore2("B Z"), do: 6 + 3

  defp getScore2(_else), do: "Error"
end
