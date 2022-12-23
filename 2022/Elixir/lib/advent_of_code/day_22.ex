defmodule AdventOfCode.Day22 do
  def part1(args) do
    findPath(args, &plain_wrap/4)
  end

  def part2(args) do
    findPath(args, &cube_wrap/4)
  end

  def findPath(args, wrapper) do
    [map, moves] = parse_input(args)

    start = Map.filter(map, fn {[_x,y], v} -> v == "." and y == 1 end) |> Enum.min() |> elem(0)
    [[last_x, last_y], last_facing, path] = moves
    |> Enum.reduce([start, :right, Map.new([{start, :right}])], fn move, [position, facing, path] ->
      [newPos, newFacing] = move(position, facing, move, map, wrapper)
      [newPos, newFacing, Map.put(path, newPos, newFacing)]
    end)

    print_path(map, path)
    calculateSolution(last_x, last_y, last_facing)
  end

  def calculateSolution(last_x, last_y, last_facing) do
    facing_value = case last_facing do
      :up ->  3
      :down -> 1
      :left -> 2
      :right -> 0
    end
    last_y * 1000 + 4 * last_x + facing_value

  end

  def print_path(map, path) do
    composite_map = Map.merge(map, path, fn _k, _vm, vp -> case vp do
      :up ->  "^"
      :down -> "v"
      :left -> "<"
      :right -> ">"
    end end)

    [xMax, yMax] = Map.keys(map) |> Enum.max()

    1..yMax
    |> Enum.map(&Enum.filter(composite_map, fn {[_x, y], _v} -> y == &1 end) |> Enum.sort() |> Enum.map(fn {_k, v} -> v end))
    |> Enum.intersperse(["\n"])
    |> IO.puts()
  end

  def parse_input(args) do
    [rawmap, rawmoves] = args |> String.split("\n\n")

    map =
      rawmap
      |> String.split("\n")
      |> Enum.with_index(fn line, row ->
        Enum.with_index(String.codepoints(line), fn letter, col -> [[col, row], letter] end)
      end)
      |> List.flatten()
      |> Enum.chunk_every(3)
      |> Enum.reduce(Map.new(), fn [x, y, k], map -> Map.put(map, [x+1, y+1], k) end)

    moves =
      rawmoves
      |> String.trim()
      |> String.codepoints()
      |> Enum.chunk_by(&Regex.match?(~r/\d/, &1))
      |> Enum.map(&Enum.join/1)
      |> Enum.map(fn x ->
        if String.contains?(x, ["L", "R"]), do: String.to_atom(x), else: String.to_integer(x)
      end)
      [map, moves]
  end

  def move(position, facing, :L, _map, _wrapper), do: [position, turnL(facing)]
  def move(position, facing, :R, _map, _wrapper), do: [position, turnL(turnL(turnL(facing)))]

  def move(position, facing, 0, _map, _wrapper), do: [position, facing]
  def move(position, facing, steps, map, wrapper) do
    [x, y] = position
    nextMove = case facing do
      :up ->  [x, y-1]
      :down -> [x, y+1]
      :left -> [x-1, y]
      :right -> [x+1, y]
    end

    case Map.get(map, nextMove, " ") do
      "." -> move(nextMove, facing, steps - 1, map, wrapper)
      "#" -> [position, facing]
      " " -> wrapper.(position, facing, steps, map)
    end
  end

  def plain_wrap(position, facing, steps, map) do
    [former_x, former_y] = position
    potential_step = case facing do
      :up ->  Map.filter(map, fn {[x,_y], v} -> v != " " and x == former_x end) |> Enum.max() |> elem(0)
      :down -> Map.filter(map, fn {[x,_y], v} -> v != " " and x == former_x end) |> Enum.min() |> elem(0)
      :left -> Map.filter(map, fn {[_x,y], v} -> v != " " and y == former_y end) |> Enum.max() |> elem(0)
      :right -> Map.filter(map, fn {[_x,y], v} -> v != " " and y == former_y end) |> Enum.min() |> elem(0)
    end

    case Map.get(map, potential_step, " ") do
      "." -> move(potential_step, facing, steps - 1, map, &plain_wrap/4)
      "#" -> [position, facing]
      " " -> :error
    end

  end

  def cube_wrap(position, facing, steps, map) do
    [former_x, former_y] = position
    rel_x = Integer.mod((former_x-1), 50) + 1
    rel_y = Integer.mod((former_y-1), 50) + 1

    # determine wrap regions xy
    # -- 10 20
    # -- 11 --
    # 02 12 --
    # 03 -- --

    x_wrap = div((former_x-1), 50)
    y_wrap = div((former_y-1), 50)


    [potential_step, potential_facing] = cond do
      facing == :up    and x_wrap == 0 and y_wrap == 2 -> [[51, 50 + rel_x], :right]
      facing == :up    and x_wrap == 1 and y_wrap == 0 -> [[1, 150 + rel_x], :right]
      facing == :up    and x_wrap == 2 and y_wrap == 0 -> [[rel_x, 200], :up]

      facing == :down  and x_wrap == 0 and y_wrap == 3 -> [[100 + rel_x, 1], :down]
      facing == :down  and x_wrap == 1 and y_wrap == 2 -> [[50, 150 + rel_x], :left]
      facing == :down  and x_wrap == 2 and y_wrap == 0 -> [[100, 50 + rel_x], :left]

      facing == :left  and x_wrap == 1 and y_wrap == 0 -> [[1, 151 - rel_y], :right]
      facing == :left  and x_wrap == 1 and y_wrap == 1 -> [[rel_y, 101], :down]
      facing == :left  and x_wrap == 0 and y_wrap == 2 -> [[51, 51 - rel_y], :right]
      facing == :left  and x_wrap == 0 and y_wrap == 3 -> [[50 + rel_y, 1], :down]

      facing == :right and x_wrap == 2 and y_wrap == 0 -> [[100, 151 - rel_y], :left]
      facing == :right and x_wrap == 1 and y_wrap == 1 -> [[100 + rel_y, 50], :up]
      facing == :right and x_wrap == 1 and y_wrap == 2 -> [[150, 51  - rel_y], :left]
      facing == :right and x_wrap == 0 and y_wrap == 3 -> [[50 + rel_y, 150], :up]

      true -> raise "case not covered!"

    end

    case Map.get(map, potential_step, " ") do
      "." -> move(potential_step, potential_facing, steps - 1, map, &plain_wrap/4)
      "#" -> [position, facing]
      " " -> [x, y] = potential_step
        raise "not implemented: potential step: x: #{x} y: #{y}"
    end

  end


  def turnL(:up), do: :left
  def turnL(:left), do: :down
  def turnL(:down), do: :right
  def turnL(:right), do: :up



end
