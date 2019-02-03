package main

import "fmt"

type song struct {

  name string
  artist string
  previous *song
  next *song

}

type playlist struct{
  name string
  head *song
  tail *song
  nowPlaying *song

}

func createPlaylist(name string) *playlist {
  return &playlist{
    name : name,
  }
}

func (p *playlist) addSong(name, author string) error {

  s := &song{
    name: name,
    artist: author,
  }
  
  if p.head == nil{
    p.head = s
  } else{
    currentNode := p.tail
    currentNode.next = s
    s.previous = p.tail
  }
  
  p.tail =s
  
  return nil

}

func (p *playlist) showAllSongs() error {
  currentNode := p.head
  fmt.Println("%+v\n", *currentNode)
  for currentNode.next != nil{
    currentNode = currentNode.next
    fmt.Println("%+v\n", *currentNode)
  }
  return nil
}

func (p *playlist) startPlaying() *song{
  p.nowPlaying = p.head
  return p.nowPlaying
}

func (p *playlist) nextSong() *song{
  p.nowPlaying = p.nowPlaying.next
  return p.nowPlaying
}

func (p *playlist) previousSong() *song{
  p.nowPlaying = p.nowPlaying.previous
  return p.nowPlaying
}

func main() {
	playlistName := "myplaylist"
	myPlaylist := createPlaylist(playlistName)
	fmt.Println("Created playlist")
	fmt.Println()

	fmt.Print("Adding songs to the playlist...\n\n")
	myPlaylist.addSong("Ophelia", "The Lumineers")
	myPlaylist.addSong("Shape of you", "Ed Sheeran")
	myPlaylist.addSong("Stubborn Love", "The Lumineers")
	myPlaylist.addSong("Feels", "Calvin Harris")
	fmt.Println("Showing all songs in playlist...")
	myPlaylist.showAllSongs()
	fmt.Println()

	myPlaylist.startPlaying()
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.previousSong()
  	fmt.Println("Changing previous song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.previousSong()
	fmt.Println("Changing previous song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
}
