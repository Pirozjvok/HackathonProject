using Microsoft.AspNetCore.Mvc;
using System.Collections.Concurrent;
using System.Text.Json;
using System.Text.RegularExpressions;

namespace Transcription
{
    public class Program
    {
        public static ConcurrentDictionary<int, string> keyValuePairs = new ConcurrentDictionary<int, string>();

        public static void Main(string[] args)
        {
            HashSet<string> names = GetNames();
            HashSet<string> surnames = GetSurNames();
            HashSet<string> midnames = GetMidNames();

            var builder = WebApplication.CreateBuilder(args);

            string aiUrl = builder.Configuration["AiUrl"];
            string audioStorage = builder.Configuration["AudioStorage"];

            var app = builder.Build();

            app.MapGet("/transcription/{id:int}", async ([FromRoute]int id) => { 
                HttpClient client = new HttpClient();
                client.Timeout = Timeout.InfiniteTimeSpan;
                Console.WriteLine("Downloading audio...");
                var response = await client.GetAsync(audioStorage + $"/{id}", HttpCompletionOption.ResponseHeadersRead);
                Console.WriteLine("Response length: {0}", response.Content.Headers.ContentLength);
                MultipartFormDataContent content = new MultipartFormDataContent();
                StreamContent asdsad = new StreamContent(await response.Content.ReadAsStreamAsync());
                content.Add(asdsad, "audio_file", $"{id}.mp3");
                HttpRequestMessage requestMessage = new HttpRequestMessage(HttpMethod.Post, aiUrl + "?encode=true&output=json&language=ru");
                requestMessage.Headers.Add("accept", "application/json");
                requestMessage.Content = content;
                var asdssa = await client.SendAsync(requestMessage);
                string json = await asdssa.Content.ReadAsStringAsync();

                JsonSerializerOptions options = new()
                {
                    PropertyNameCaseInsensitive = true
                };
                Response myDeserializedClass = JsonSerializer.Deserialize<Response>(json, options)!;

                Chat chat = new Chat();

                chat.Messages = myDeserializedClass.segments.Select(x =>
                {
                    return new Message()
                    {
                        FromOperator = false,
                        Text = x.text
                    };
                }).ToList();

                string? name = GetName(chat.Messages, names);
                string? surname = GetName(chat.Messages, surnames);
                string? midname = GetName(chat.Messages, midnames);

                chat.Name = name;
                chat.SurName = surname;
                chat.MidName = midname;

                return chat;
            });
            app.Run();
        }

        public static HashSet<string> GetNames()
        {
            HashSet<string> names = new HashSet<string>();
            string[] lines = File.ReadAllLines("names_table.jsonl");
            foreach (string line in lines)
            {
                string name = JsonDocument.Parse(line).RootElement.GetProperty("text").GetString()!.ToLower();
                names.Add(name);
            }
            return names;
        }

        public static HashSet<string> GetSurNames()
        {
            HashSet<string> names = new HashSet<string>();
            string[] lines = File.ReadAllLines("surnames_table.jsonl");
            foreach (string line in lines)
            {
                names.Add(JsonDocument.Parse(line).RootElement.GetProperty("text").GetString()!.ToLower());
            }
            return names;
        }

        public static HashSet<string> GetMidNames()
        {
            HashSet<string> names = new HashSet<string>();
            string[] lines = File.ReadAllLines("midnames_table.jsonl");
            foreach (string line in lines)
            {
                names.Add(JsonDocument.Parse(line).RootElement.GetProperty("text").GetString()!.ToLower());
            }
            return names;
        }

        public static string? GetName(IEnumerable<Message> messages, HashSet<string> names)
        {
            Regex regex = new Regex(@"(\w+)");

            string? current_name = null;

            foreach (var msg in messages)
            {
                string[] words = regex.Matches(msg.Text).Select(x => x.Value).ToArray();

                foreach (string word in words)
                {
                    var word_lower = word.ToLower();
                    if (word.Length >= 1 && char.IsUpper(word[0]) && names.Contains(word_lower))
                    {
                        if (current_name == null)
                        {
                            current_name = word;
                        } 
                        else if (word.Length > current_name.Length)
                        {
                            current_name = word;
                        }
                    }
                }
            }
            return current_name;
        }
    }

    // Root myDeserializedClass = JsonConvert.DeserializeObject<Root>(myJsonResponse);
    public class Response
    {
        #pragma warning disable
        public string text { get; set; }
        public List<Segment> segments { get; set; }
        public string language { get; set; }
    }

    public class Segment
    {
        public int id { get; set; }
        public int seek { get; set; }
        public double start { get; set; }
        public double end { get; set; }
        public string text { get; set; }
        public List<int> tokens { get; set; }
        public double temperature { get; set; }
        public double avg_logprob { get; set; }
        public double compression_ratio { get; set; }
        public double no_speech_prob { get; set; }
    }

    public class Chat
    {
        public List<Message> Messages { get; set; }

        public string? Name { get; set; }

        public string? SurName { get; set; }

        public string? MidName { get; set; }

        public Chat()
        {
            Messages = new();
        }
    }

    public class Message
    {
        public string Text { get; set; }

        public bool FromOperator { get; set; }
    }
}