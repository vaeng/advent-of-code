defmodule AdventOfCode.Day07 do
  @spec parseCommand(String.t()) :: list()
  def parseCommand(rawStr) do
    cond do
      String.starts_with?(rawStr, "$ ls") ->
        [:ls]

      String.starts_with?(rawStr, "$ cd") ->
        [:cd, String.trim_leading(rawStr, "$ cd ")]

      String.starts_with?(rawStr, "dir") ->
        [:dir, String.trim_leading(rawStr, "dir ")]

      String.match?(rawStr, ~r/^\d+ .+$/) ->
        [size, name] = String.split(rawStr, " ")
        [:file, String.to_integer(size), name]

      true ->
        [:unknown, rawStr]
    end
  end

  def executeCommand([:cd, "/"], [_path, rootNode, folders]),
    do: [["/"], rootNode, folders]

  def executeCommand([:cd, ".."], [["/"], rootNode, folders]),
    do: [["/"], rootNode, folders]

  def executeCommand([:cd, ".."], [path, rootNode, folders]),
    do: [Enum.drop(path, -1), rootNode, folders]

  def executeCommand([:cd, folder], [path, rootNode, folders]),
    do: [path ++ [folder], rootNode, MapSet.put(folders, path ++ [folder])]

  def executeCommand([:ls], [path, rootNode, folders]),
    do: [path, rootNode, folders]

  def executeCommand([:file, size, name], [path, rootNode, folders]),
    do: [path, update_in(rootNode, path, &Map.put(&1, name, size)), folders]

  def executeCommand([:dir, name], [path, rootNode, folders]),
    do: [path, update_in(rootNode, path, &Map.put(&1, name, %{})), folders]

  def part1(args) do
    root = %{"/" => %{}}

    [_, tree, visitedPaths] =
      args
      |> String.trim()
      |> String.split("\n")
      |> Enum.map(&parseCommand/1)
      |> Enum.reduce(["/", root, MapSet.new([["/"]])], fn x, acc -> executeCommand(x, acc) end)

    visitedPaths
    |> Enum.map(&(get_in(tree, &1) |> getSize()))
    |> Enum.filter(&(&1 <= 100_000))
    |> Enum.sum()
  end

  def part2(args) do
    root = %{"/" => %{}}

    [_, tree, visitedPaths] =
      args
      |> String.trim()
      |> String.split("\n")
      |> Enum.map(&parseCommand/1)
      |> Enum.reduce(["/", root, MapSet.new([["/"]])], fn x, acc -> executeCommand(x, acc) end)

    otherDirectories = MapSet.delete(visitedPaths, [["/"]])

    space_available = 70_000_000 - getSize(get_in(tree, ["/"]))
    additional_space_needed = 30_000_000 - space_available

    directorySizes =
      otherDirectories
      |> Enum.map(&(get_in(tree, &1) |> getSize()))
      |> Enum.sort()
      |> Enum.filter(&(&1 >= additional_space_needed))
      |> hd
  end

  @spec getSize(list) :: number
  @doc """
  Returns the size of a filetree in form of a Map.
  """
  def getSize(element) when is_integer(element), do: element

  def getSize(tuple) when is_tuple(tuple) and tuple_size(tuple) == 2 do
    {_, map} = tuple
    getSize(map)
  end

  def getSize(map) when is_map(map) do
    map |> Map.to_list() |> List.flatten() |> Enum.map(&getSize/1) |> Enum.sum()
  end
end
