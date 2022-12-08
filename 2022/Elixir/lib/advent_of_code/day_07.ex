defmodule AdventOfCode.Day07 do
  defmodule Folder do
    @enforce_keys [:name]
    defstruct [:name, files: %{}, folders: %{}, type: :dir]
  end

  defmodule File do
    @enforce_keys [:name, :size]
    defstruct [:name, :size, type: :file]
  end

  @doc"""
  defimpl Inspect, for: Folder do
    def inspect(folder, _opts) do
      folderText = "" # folder.folders.keys()
      fileText = "|- a file\n"
      "-- {folder.name}\n" <> folderText <> fileText
    end
  end

  defimpl Inspect, for: File do
    def inspect(file, _opts) do
      "-- {file.size} {file.name}\n"
    end
  end
  """
  defprotocol FileTree do
    def size(value)
  end

  defimpl FileTree, for: Folder do
    def size(dir), do:
      dir.content
      |> Enum.map(&FileTree.size/1)
      |> Enum.sum()
  end

  defimpl FileTree, for: File do
    def size(file), do: file.size
  end

  @spec parseCommand(String.t()) :: list()
  def parseCommand(rawStr) do
    cond do
      String.starts_with?(rawStr, "$ ls") -> [:ls]
      String.starts_with?(rawStr, "$ cd") -> [:cd, String.trim_leading(rawStr, "$ cd ")]
      String.starts_with?(rawStr, "dir") -> [:dir, String.trim_leading(rawStr, "dir ")]
      String.match?(rawStr, ~r/^\d+ .+$/) ->
        [size, name] = String.split(rawStr, " ")
        [:file, String.to_integer(size), name]
      true -> [:unknown, rawStr]
    end
  end

  @spec executeCommand(list(), list()) :: %Folder{}
  def executeCommand([:cd, "/"], [_path, rootNode]), do: [["/"], rootNode]
  def executeCommand([:cd, ".."], [["/"], rootNode]), do: [["/"], rootNode]
  def executeCommand([:cd, ".."], [path, rootNode]), do: [Enum.drop(path, -1), rootNode]
  def executeCommand([:cd, folder], [path, rootNode]) do
     # nextNode = Enum.find(currentNode.content, fn x -> x.type == :dir && x.name == folder  end)
     [ path ++ [folder], rootNode]
  end
  def executeCommand([:ls], [path, rootNode]), do: [path, rootNode]
  def executeCommand([:file, size, name], [path, rootNode])
    do
      newFile = %File{name: name, size: size}
      newTree = addItemToContent(path, rootNode, newFile)
      [path, newTree]
    end
  def executeCommand([:dir, name], [path, rootNode])
    do
      newDir = %Folder{name: name}
      newTree = addItemToContent(path, rootNode, newDir)
      [path, newTree]
    end

  @spec addItemToContent(list(), %Folder{}, %Folder{} | %File{}) :: %Folder{}
  def addItemToContent(path, root, item) do
    access_path = transformToAccessPath(path)
    dir = if access_path == [], do: root, else: get_in(root, access_path)
    IO.inspect(dir)
    if duplicate?(dir, item)
      do
        root
      else
        # get_and_update_in(root, [Access.key(:files), Access.key(:abaa), Access.key(:size)], &{&1, &1 + 1000})
        #get_and_update_in(root, path, fn currentFolder -> %{currentFolder | content: Map.update(dir.content, item.name, item)} end)
        if access_path != []
          do
            update_in(root, access_path, &(Map.put(&1, item.name,  item)))
          else
            key = if item.type == :dir, do: :folders, else: :files
            update(root, )
          end
      end
    end

  def transformToAccessPath(path) do
    path |> Enum.map(&Access.key/1) |> Enum.intersperse(Access.key(:folders)) |> Enum.drop(1)
  end

  @spec duplicate?(%Folder{}, %Folder{} | %File{}) :: boolean()
  def duplicate?(folder, item) do
    key = if item.type == :dir, do: :folders, else: :files
    get_in(folder, [Access.key(key)]) |> Map.keys() |> Enum.member?(item.name)
  end

  def part1(args) do
    root = %Folder{name: "/"}

    tree = args
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&parseCommand/1)
    |> Enum.reduce(["/", root], fn x, acc -> executeCommand(x, acc) end)

    IO.inspect(tree)
  end

  def part2(_args) do
  end

  @spec getSize(list) :: number
  @doc """
  Returns the size of a filetree in form of a Map.
  """
  def getSize(tree) do
    tree |> Map.to_list() |> List.flatten() |> Enum.map(&getSize/1) |> Enum.sum()
  end
end
