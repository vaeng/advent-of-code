defmodule AdventOfCode.Day17 do
  @starttime DateTime.utc_now()

  def part1(args) do
    # parse input
    directions = parseInputToStream(args)


    width = 7
    solidline = List.duplicate(:rock, width + 2)
    startLayout = %{0 => solidline}
    droplimit = 1000000000000
    startRock = :horiz
    initialFallingRock = initialPosition(startRock, 1)

    solution = directions
    |> Enum.reduce_while(
      [%{}, Map.new(), startLayout, startRock, initialFallingRock, 1, 1],
      fn {currentDir, idx}, [heights, map, layout, currentRock, fallingRock, lastLine, rocksdropped] ->
        newItem = [idx, currentRock, Map.fetch!(layout, lastLine - 1)]
        if Map.has_key?(map, newItem) do
          firstSeen = Map.get(map, newItem)
          currently = rocksdropped
          cycleLength = firstSeen - currently
          first_height = Map.get(heights, firstSeen)
          cycleHeight = lastLine - first_height
          cycles_left = div(droplimit - rocksdropped, cycleLength) - 1
          lastLine = lastLine + cycles_left * cycleHeight
          rocksdropped = rocksdropped + cycles_left * cycleLength

          {status, [layout, currentRock, movedPositions, lastLine, rocksdropped]} = handleRock(layout, fallingRock, currentDir, currentRock, lastLine, rocksdropped, droplimit)
          {status, [Map.new(), Map.new() | [layout, currentRock, movedPositions, lastLine, rocksdropped]]}

          # {:halt, [Map.fetch!(map, newItem), rocksdropped, [Map.put_new(heights, rocksdropped, lastLine), idx, layout, fallingRock, currentDir, currentRock, lastLine, rocksdropped]]}
        else
         {status, [layout, currentRock, movedPositions, lastLine, rocksdropped]} = handleRock(layout, fallingRock, currentDir, currentRock, lastLine, rocksdropped, droplimit)
         {status, [Map.put_new(heights, rocksdropped, lastLine), Map.put(map, newItem, rocksdropped) | [layout, currentRock, movedPositions, lastLine, rocksdropped]]}
        end
      end
    )

    cond do
      is_integer(solution) -> solution
      true ->
        [firstSeen, againSeen, [heights, idx, layout, fallingRock, currentDir, currentRock, lastLine, rocksdropped]] = solution
        # time until entering the loop:
        init = 1000000000000 - firstSeen
        loop_length = againSeen - firstSeen
        loop_counter = div(init, loop_length)
        rocks_left = Integer.mod(init, loop_length)

        height_at_loop_start = Map.get(heights, firstSeen)
        height_at_loop_end = Map.get(heights, againSeen)
        loop_height = height_at_loop_end - height_at_loop_start
        rest_of_the_rocks = Map.get(heights, firstSeen + rocks_left)
        result = height_at_loop_start + loop_height * loop_counter + rest_of_the_rocks
        1514285714288 - result
    end
  end

  def updateStatus(rocksdropped,droplimit) do
    if Integer.mod(rocksdropped, 1000_000) == 0 do
      percentage = Float.round(100*rocksdropped/droplimit, 1)
      time_elapsed = DateTime.diff(DateTime.utc_now(), @starttime)
      forecast = time_elapsed / rocksdropped * droplimit
      hours_elapsed = Float.round(time_elapsed/ 60 / 60 )
      hours_forecast = Float.round(forecast/ 60 / 60)
     IO.puts(IO.ANSI.clear_line() <> "Dropped Rock #{rocksdropped} of #{droplimit} (#{percentage} %) Elapsed time (#{hours_elapsed} hours of #{hours_forecast})")
    end
  end

  def parseInputToStream(args) do
    args
      |> String.trim()
      |> String.codepoints()
      |> Enum.with_index()
      |> Stream.map(fn {dir, idx} -> if dir == ">", do: {:right, idx}, else: {:left, idx} end)
      |> Stream.cycle()
  end

  def handleRock(layout, fallingRock, :down, currentRock, lastLine, rocksdropped, droplimit) do

    if rocksdropped == droplimit do
      # inspectLayout(layout, 0, lastLine)
      {:halt, lastLine}
    else
      movedPositions = move(fallingRock, :down)

      case willCollide?(layout, movedPositions) do
        true ->
          memoryLimit = 100
          updatedLayout = place(layout, fallingRock)

          newLastline = Enum.max(Map.keys(updatedLayout))
          trimFrom = newLastline - 2 * memoryLimit
          trimTo = newLastline - memoryLimit

          trimmedLayout =
            Enum.reduce(trimFrom..trimTo, updatedLayout, fn x, acc ->
              Map.delete(acc, x)
            end)

          nextRock = getNextRock(currentRock)
          nextFallingRockPosition = initialPosition(nextRock, newLastline + 1)
          # fallingRock}

          {:cont,
           [trimmedLayout, nextRock, nextFallingRockPosition, newLastline, rocksdropped + 1]}

        # movedPositions}
        false ->
          {:cont, [layout, currentRock, movedPositions, lastLine, rocksdropped]}
      end
    end
  end

  def handleRock(layout, fallingRock, direction, currentRock, lastLine, rocksdropped, droplimit) do
    movedPositions = move(fallingRock, direction)

    case willCollide?(layout, movedPositions) do
      true -> handleRock(layout, fallingRock, :down, currentRock, lastLine, rocksdropped, droplimit)
      false -> handleRock(layout, movedPositions, :down, currentRock, lastLine, rocksdropped, droplimit)
    end
  end

  def getNextRock(:plus), do: :angle
  def getNextRock(:angle), do: :vert
  def getNextRock(:vert), do: :block
  def getNextRock(:block), do: :horiz
  def getNextRock(:horiz), do: :plus

  def willCollide?(layout, fallingRock) do
    Enum.any?(fallingRock, fn [x, y] ->
      line = Map.get(layout, y, [:rock] ++ List.duplicate(:air, 7) ++ [:rock])
      Enum.at(line, x) == :rock
    end)
  end

  def inspectLayout(layout, from, to, movingRocks \\ []) do
    movingLayout = place(layout, movingRocks, :movingRock)

    to..from//-1
    |> Enum.each(
      &(Map.get(movingLayout, &1, [:rock] ++ List.duplicate(:air, 7) ++ [:rock])
        |> Enum.map(fn x ->
          case x do
            :rock -> "#"
            :air -> "."
            :movingRock -> "@"
          end
        end)
        |> Enum.join()
        |> IO.puts())
    )

    IO.puts("\n")
    layout
  end

  def place(layout, restingRocks, replacement \\ :rock) do
    Enum.reduce(restingRocks, layout, fn [x, y], acc ->
      line =
        Map.get(layout, y, [:rock] ++ List.duplicate(:air, 7) ++ [:rock])
        |> List.replace_at(x, replacement)

      Map.update(acc, y, line, &List.replace_at(&1, x, replacement))
    end)
  end

  def move(fallingRock, :down), do: Enum.map(fallingRock, fn [x, y] -> [x, y - 1] end)
  def move(fallingRock, :right), do: Enum.map(fallingRock, fn [x, y] -> [x + 1, y] end)
  def move(fallingRock, :left), do: Enum.map(fallingRock, fn [x, y] -> [x - 1, y] end)

  def initialPosition(:horiz, lastLine) do
    [[3, lastLine + 3], [4, lastLine + 3], [5, lastLine + 3], [6, lastLine + 3]]
  end

  def initialPosition(:plus, lastLine) do
    [
      [3, lastLine + 4],
      [4, lastLine + 4],
      [4, lastLine + 5],
      [4, lastLine + 3],
      [5, lastLine + 4]
    ]
  end

  def initialPosition(:angle, lastLine) do
    [
      [3, lastLine + 3],
      [4, lastLine + 3],
      [5, lastLine + 3],
      [5, lastLine + 4],
      [5, lastLine + 5]
    ]
  end

  def initialPosition(:vert, lastLine) do
    [[3, lastLine + 3], [3, lastLine + 4], [3, lastLine + 5], [3, lastLine + 6]]
  end

  def initialPosition(:block, lastLine) do
    [[3, lastLine + 3], [4, lastLine + 3], [3, lastLine + 4], [4, lastLine + 4]]
  end

  def part2(_args) do
  end
end
