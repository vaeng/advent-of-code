defmodule AdventOfCode.Day01 do
  def part1(args) do
    args
    |> String.to_charlist()
    |> Enum.reduce(0, fn x, acc -> if( x == 41, do: acc-1, else: acc+1) end)
  end

  def part2(args) do
    args
    |> String.to_charlist()
    |> Enum.reduce_while([0, 0], &updateLevel(&1, &2))
    |> hd
  end

  defp updateLevel(_, [step, level]) when level == -1, do: {:halt, [step, level]}
  defp updateLevel(x, [step, level]) when x == 40, do: {:cont, [step+1, level+1]}
  defp updateLevel(x, [step, level]) when x == 41, do: {:cont, [step+1, level-1]}
end
