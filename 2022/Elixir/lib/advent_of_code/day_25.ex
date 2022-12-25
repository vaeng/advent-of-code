defmodule AdventOfCode.Day25 do
  def part1(args) do
    args
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&convert_from_snafu/1)
    |> Enum.sum()
    |> (then &convert_to_snafu/1)
  end

  def part2(_args) do
  end

  def convert_letter(c) do
    case c do
      "=" -> "0"
      "-" -> "1"
      "0" -> "2"
      "1" -> "3"
      "2" -> "4"
    end
  end

  def convert_digit(c) do
    case c do
      2 -> "0"
      3 -> "1"
      4 -> "2"
      1 -> "-"
      0 -> "="
    end
  end

  def convert_from_snafu(string) do
    bond = string
    |> String.trim()
    |> String.length()
    |> (then &List.duplicate("0", &1))
    |> Enum.map(&convert_letter/1)
    |> Enum.join()
    |> String.to_integer(5)
    num = string
    |> String.trim()
    |> String.codepoints()
    |> Enum.map(&convert_letter/1)
    |> Enum.join()
    |> String.to_integer(5)

    num - bond
  end

  def convert_to_snafu(number) do
    bound = Integer.undigits(List.duplicate(2, Enum.count(Integer.digits(number, 5))+1), 5)
    Integer.digits(bound - number, 5)
    |> Enum.map(fn n -> case n do
      1 -> "1"
      2 -> "0"
      3 -> "-"
      4 -> "="
      0 -> "2"
    end
    end)
    |> Enum.join()
    |> String.trim_leading("0")
  end

end
