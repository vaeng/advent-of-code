defmodule AdventOfCode.Day24 do
  def part1(args) do
    initial_map = parse_input(args)
    [{start_position, _}] = Enum.filter(map, fn {[_x,y], k} -> y == 0 and k == [] end)

    0..100
    |> Enum.reduce_while([initial_map, [start_position]], fn minute, [map, pos_expedition_position] ->

    end)

  def part2(_args) do
  end

  def progress_blizzards(map) do
    Enum.reduce(map, Map.new(), fn {pos, types}, accmap ->
      Enum.reduce(types, accmap, fn type, inner_map ->
        new_pos = move_blizzard(map, {pos, type})
        Map.update(inner_map, new_pos, [type], fn old -> [type | old] end)
      end)
    end)
  end

  def move_blizzard(map, {[x, y], type}) do
    [x_max, y_max] = Enum.max(Map.keys(map))
    new_pos = case type do
      :blizzard_right -> if x + 1 < x_max, do: [x + 1, y], else: [1, y]
      :blizzard_up -> if y - 1 > 0, do: [x, y - 1], else: [x, y_max - 1]
      :blizzard_left -> if x - 1 > 0, do: [x - 1, y], else: [x_max - 1, y]
      :blizzard_down -> if y + 1 < y_max, do: [x, y + 1], else: [x, 1]
      _ -> [x, y]
    end
    new_pos
  end

  def parse_input(args) do
    args
    |> String.split("\n")
    |> Enum.with_index(fn line, row ->
      Enum.with_index(String.codepoints(line), fn letter, col ->
      key = case letter do
        "." -> :free
        "#" -> :wall
        "<" -> :blizzard_left
        "^" -> :blizzard_up
        "v" -> :blizzard_down
        ">" -> :blizzard_right
      end
      [[col, row], key] end)
    end)
    |> List.flatten()
    |> Enum.chunk_every(3)
    |> Enum.reduce(Map.new(), fn [x, y, k], map -> unless k == :free, do: Map.put(map, [x, y], [k]), else: Map.put(map, [x, y], []) end)
  end

  def print(map) do
      [x_max, y_max] = Map.keys(map) |> Enum.max()
      0..y_max
      |> Enum.map(
        fn y -> Enum.map(0..x_max,
            fn x ->
              content = Map.get(map, [x, y], [])
              cond  do
                content == [:E] -> "E"
                content == [] -> "."
                content == [:wall] -> "#"
                content == [:blizzard_left] -> "<"
                content == [:blizzard_up] -> "^"
                content == [:blizzard_down] -> "v"
                content == [:blizzard_right] -> ">"
                true -> Integer.to_string(Enum.count(content))
              end
            end
        ) end)
      |> Enum.intersperse(["\n"])
      |> IO.puts()
  end
end
