defmodule AdventOfCode.Day17 do
  def part1(args) do
    # parse input
    directions =
      args
      |> String.trim()
      |> String.codepoints()
      |> Stream.map(fn dir -> if dir == ">", do: :right, else: :left end)
      |> Stream.cycle()

    width = 7

    startLayout = %{0 => List.duplicate(:rock, width + 2)}
    rockLimit = 2022
    startRock = :horiz
    initialFallingRock = initialPosition(startRock, 1)

    directions
    |> Enum.reduce_while(
      [startLayout, startRock, initialFallingRock, 1, 1],
      fn currentDir, [layout, currentRock, fallingRock, lastLine, rocksdropped] ->
        handleRock(layout, fallingRock, currentDir, currentRock, lastLine, rocksdropped)
      end
    )
  end

  def handleRock(layout, fallingRock, :down, currentRock, lastLine, rocksdropped) do
    if rocksdropped == 1_000_000_000_000 + 1 do
      # inspectLayout(layout, 0, lastLine)
      {:halt, lastLine}
    else
      movedPositions = move(fallingRock, :down)

      case willCollide?(layout, movedPositions) do
        true ->
          updatedLayout = place(layout, fallingRock)

          newLastline = Enum.max(Map.keys(updatedLayout))
          trimFrom = newLastline - 50
          trimTo = trimFrom + 25

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

  def handleRock(layout, fallingRock, direction, currentRock, lastLine, rocksdropped) do
    movedPositions = move(fallingRock, direction)

    case willCollide?(layout, movedPositions) do
      true -> handleRock(layout, fallingRock, :down, currentRock, lastLine, rocksdropped)
      false -> handleRock(layout, movedPositions, :down, currentRock, lastLine, rocksdropped)
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
