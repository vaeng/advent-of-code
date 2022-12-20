defmodule AdventOfCode.Day20 do
  def part1(args) do
    initialArray = args |> String.split() |> Enum.with_index(fn idx, el -> {idx, String.to_integer(el)} end)
    IO.puts("Initial arrangement:")
    IO.puts("#{initialArray |> Enum.map(&elem(&1,1)) |> Enum.join(", ")}")
    IO.puts("")

    length = Enum.count(initialArray)

    0..length-1
    |> Enum.reduce(initialArray, fn idx, arr ->
        current_pos = Enum.find_index(arr, fn {i, _el} -> i == idx end)
        current_element = Enum.at(arr, current_pos) |> elem(1)

        if current_element == 0 do
          IO.puts("0 does not move:")
          IO.puts("#{arr |> Enum.map(&elem(&1,1)) |> Enum.join(", ")}")
          IO.puts("")
          arr
        else
          raw_pos = current_pos + current_element
          limit =
          intended_pos = cond do
            0 < raw_pos and raw_pos < length -> raw_pos

            true -> raw_pos + div(raw_pos/length)
          end


          newArray = List.delete_at(arr, current_pos) |> List.insert_at(intended_pos, {idx, current_element})

          IO.puts("#{current_element} moves between #{Enum.at(arr, intended_pos) |> elem(1)} and #{Enum.at(arr, Integer.mod(intended_pos + 1, length)) |> elem(1)}:")
          IO.puts("#{newArray |> Enum.map(&elem(&1,1)) |> Enum.join(", ")}")
          IO.puts("")
          newArray
        end
      end)
  end

  def part2(_args) do
  end
end
