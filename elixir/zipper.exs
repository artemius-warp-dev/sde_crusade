defmodule Zipper do
  defstruct thread: [], fork: nil

  @type t :: %Zipper{thread: list, fork: BinTree.t() | nil}

  alias BinTree

  @doc """
  Get a zipper focused on the root node.
  """
  @spec from_tree(BinTree.t()) :: Zipper.t()
  def from_tree(bin_tree) do
    %Zipper{
      thread: [],
      fork: bin_tree
    }
  end

  @doc """
  Get the complete tree from a zipper.
  """
  @spec to_tree(Zipper.t()) :: BinTree.t()
  def to_tree(%Zipper{thread: [], fork: fork}), do: fork

  def to_tree(zipper) do
    up(zipper)
    |> to_tree()
  end

  @doc """
  Get the value of the focus node.
  """
  @spec value(Zipper.t()) :: any
  def value(%Zipper{fork: %BinTree{value: value}}), do: value

  @doc """
  Get the left child of the focus node, if any.
  """
  @spec left(Zipper.t()) :: Zipper.t() | nil
  def left(zipper) do
    %{
      zipper
      | thread: [{:left, zipper.fork.value, zipper.fork.right} | zipper.thread],
        fork: zipper.fork.left
    }
    |> case do
      %Zipper{fork: nil} -> nil
      zipper -> zipper
    end
  end

  @doc """
  Get the right child of the focus node, if any.
  """
  @spec right(Zipper.t()) :: Zipper.t() | nil
  def right(zipper) do
    %{
      zipper
      | thread: [{:right, zipper.fork.value, zipper.fork.left} | zipper.thread],
        fork: zipper.fork.right
    }
    |> case do
      %Zipper{fork: nil} -> nil
      zipper -> zipper
    end
  end

  @doc """
  Get the parent of the focus node, if any.
  """
  @spec up(Zipper.t()) :: Zipper.t() | nil
  def up(%Zipper{thread: []}), do: nil

  def up(%Zipper{thread: [{:left, value, right} | thread], fork: left}),
    do: %Zipper{thread: thread, fork: %BinTree{value: value, left: left, right: right}}

  def up(%Zipper{thread: [{:right, value, left} | thread], fork: right}),
    do: %Zipper{thread: thread, fork: %BinTree{value: value, left: left, right: right}}

  @doc """
  Set the value of the focus node.
  """
  @spec set_value(Zipper.t(), any) :: Zipper.t()
  def set_value(%Zipper{fork: nil} = zipper, value) do
    %{zipper | fork: %BinTree{value: value}}
  end

  def set_value(%Zipper{fork: %BinTree{}} = zipper, value) do
    %{zipper | fork: %{zipper.fork | value: value}}
  end

  @doc """
  Replace the left child tree of the focus node.
  """
  @spec set_left(Zipper.t(), BinTree.t() | nil) :: Zipper.t()
  def set_left(zipper, left) do
    %{zipper | fork: %{zipper.fork | left: left}}
  end

  @doc """
  Replace the right child tree of the focus node.
  """
  @spec set_right(Zipper.t(), BinTree.t() | nil) :: Zipper.t()
  def set_right(zipper, right) do
    %{zipper | fork: %{zipper.fork | right: right}}
  end
end

defmodule BinTree do
  @moduledoc """
  A node in a binary tree.

  `value` is the value of a node.
  `left` is the left subtree (nil if no subtree).
  `right` is the right subtree (nil if no subtree).
  """

  @type t :: %BinTree{value: any, left: t() | nil, right: t() | nil}

  defstruct [:value, :left, :right]
end

defimpl Inspect, for: BinTree do
  import Inspect.Algebra

  # A custom inspect instance purely for the tests, this makes error messages
  # much more readable.
  #
  # For example:
  #
  # - %BinTree{value: 3, left: nil, right: nil}
  #   becomes
  #   (3::)
  #
  # - %BinTree{value: 3, left: nil, right: %BinTree{value: 5}}
  #   becomes
  #   (3::(5::))
  #
  # - %BinTree{value: 3, left: %BinTree{value: 1}, right: %BinTree{value: 5}}
  #   becomes
  #   (3:(1::):(5::))

  def inspect(%BinTree{value: value, left: left, right: right}, opts) do
    concat([
      "(",
      to_doc(value, opts),
      ":",
      if(left, do: to_doc(left, opts), else: ""),
      ":",
      if(right, do: to_doc(right, opts), else: ""),
      ")"
    ])
  end
end
