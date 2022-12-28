defmodule AdventOfCode.Day24 do
  def part1(args) do
    initial_map = parse_input(args)
    [_, y_max] = Enum.max(Map.keys(initial_map))
    [{start_position, _}] = Enum.filter(initial_map, fn {[_x, y], k} -> y == 0 and k == [] end)
    make_run(initial_map, 0, y_max, start_position) |> hd
  end

  def part2(args) do
    initial_map = parse_input(args)
    [_, y_max] = Enum.max(Map.keys(initial_map))
    [{start_position, _}] = Enum.filter(initial_map, fn {[_x, y], k} -> y == 0 and k == [] end)
    [{goal_position, _}] = Enum.filter(initial_map, fn {[_x, y], k} -> y == y_max and k == [] end)
    [to_goal, temp_map] = make_run(initial_map, 0, y_max, start_position)
    [back_to_start, temp_map] = make_run(temp_map, to_goal, 0, goal_position)
    [back_to_goal, _] = make_run(temp_map, back_to_start, y_max, start_position)
    back_to_goal
  end

  def make_run(start_map, start_time, goal_line, start_position) do
    [x_max, y_max] = Enum.max(Map.keys(start_map))

    start_time..100_000
    |> Enum.reduce_while([start_map, [start_position]], fn minute, [map, possible_expedition_positions] ->

      next_map = progress_blizzards(map)
      next_moves = possible_expedition_positions
                           |> Enum.reduce([], fn pos, lst -> get_possible_moves(pos) ++ lst end)
                           |> Enum.uniq()
                           |> Enum.filter(fn [x, y] -> x >= 0 and y >= 0 and x <= x_max and y <= y_max and Map.get(next_map, [x, y], []) == [] end)
      if finished?(next_moves, goal_line) do
        {:halt, [minute + 1, next_map]}
      else
        {:cont, [next_map, next_moves]}
      end

    end)
  end

  def finished?(moves, goal_line) do
    Enum.any?(moves, fn [_x, y] -> y == goal_line end)
  end

  def get_possible_moves([x, y]) do
    [[x, y], [x, y + 1], [x, y - 1], [x + 1, y], [x - 1, y]]
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
