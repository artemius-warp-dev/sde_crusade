defmodule Newsletter do
  def read_emails(path) do
    File.read!(path)
    |> String.split("\n")
    |> Enum.filter(&(&1 != ""))
  end

  def open_log(path) do
    File.open!(path, [:write])
  end

  def log_sent_email(pid, email) do
    IO.puts(pid, email)
  end

  def close_log(pid) do
    File.close(pid)
  end

  def send_newsletter(emails_path, log_path, send_fun) do
    log_pid = open_log(log_path)
    emails_path
    |> read_emails()
    |> Enum.filter(&(not (&1 in read_emails(log_path))))
    |> Enum.map(&(log(send_fun.(&1), log_pid, &1)))
    close_log(log_pid)
  end
