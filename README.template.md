# yeetgif

Composable GIF effects CLI, with reasonable defaults. Made for custom Slack/Discord emoji :)

![terminal](doc/terminal.gif)

- [Get it](#get-it)
    - [Alternative 1: `go get`](#alternative-1-go-get)
    - [Alternative 2: just download the binary](#alternative-2-just-download-the-binary)
    - [Alternative 3: docker](#alternative-3-docker)
- [Use it](#use-it)
- [Hall of Fame](#hall-of-fame)
- [Usage](#usage)
    - [Conventions & tips](#conventions--tips)
    - [roll](#roll)
    - [wobble](#wobble)
    - [pulse](#pulse)
    - [zoom](#zoom)
    - [shake](#shake)
    - [woke](#woke)
    - [fried](#fried)
    - [hue](#hue)
    - [tint](#tint)
    - [resize](#resize)
    - [crop](#crop)
    - [optimize](#optimize)
    - [compose](#compose)
    - [crowd](#crowd)
    - [erase](#erase)
    - [chop](#chop)
    - [text](#text)
    - [emoji](#emoji)
    - [rain](#rain)
    - [cat](#cat)
    - [meta](#meta)
- [Licensing](#licensing)

## Get it

### Alternative 1: `go get`

```sh
go get -u github.com/sgreben/yeetgif/cmd/gif
```

### Alternative 2: just download the binary

Either from [the releases page](https://github.com/sgreben/yeetgif/releases/latest), or from the shell:

```sh
# Linux
curl -L https://github.com/sgreben/yeetgif/releases/download/${VERSION}/gif_${VERSION}_linux_x86_64.tar.gz | tar xz

# OS X
curl -L https://github.com/sgreben/yeetgif/releases/download/${VERSION}/gif_${VERSION}_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/yeetgif/releases/download/${VERSION}/gif_${VERSION}_windows_x86_64.zip
unzip gif_${VERSION}_windows_x86_64.zip
```

**NOTE**: To use the `optimize` command, you'll also need the [`giflossy`](https://github.com/kornelski/giflossy) fork of `gifsicle` installed:

```sh
brew install giflossy
```

You'll likely also want to have the binary in your `$PATH`. You can achieve this by adding this to your .bashrc (or .zshrc, ...):

```sh
export PATH=<directory-containing-the-gif-binary>:$PATH
```

### Alternative 3: docker

```sh
docker pull quay.io/sergey_grebenshchikov/yeetgif
docker tag quay.io/sergey_grebenshchikov/yeetgif gif # (optional)
```

## Use it

```sh
<doc/yeet.png gif fried | gif wobble  >doc/yeet.gif
```
![before](doc/yeet.png)
![after](doc/yeet.gif)


```sh
gif emoji aubergine | gif wobble >doc/eggplant_wobble.gif
```
![before](doc/eggplant.png)
![after](doc/eggplant_wobble.gif)

## Hall of Fame

Post a GIF made using yeetgif with either the

- [`#yeetgif` Twitter hashtag](https://twitter.com/hashtag/yeetgif?f=tweets)
- and/or the [`#yeetgif` Giphy hashtag](https://giphy.com/search/yeetgif-stickers)
- and/or the [`#yeetgif` Imgur hashtag](https://imgur.com/t/yeetgif)

~~Best~~ Most utterly demented ones end up below!

> No entries yet. Be the first :)

## Usage

```text
${USAGE}
```

### Conventions & tips

- To find out how a given example was made, try running `gif meta show` on it (e.g. `<yeet.gif gif meta show -p` will print the shell pipe of gif effects used to create `yeet.gif`).
- Use the `--raw` (`-r`) option for intermediate pipe steps -- this is faster than re-encoding as GIF every time. Also it's lossless.
- Options with bracketed default values (e.g. `--noise` of [`fried`](#fried)) can take comma-separated values - the points will be spread over the animation length, with intermediate values linearly interpolated.
- To figure out what a parameter does, try out some values around its default value, as well as much larger/smaller ones.
- To reduce GIF size, try specifying a smaller number of duplicates for static images (e.g. `gif -n 20`), `gif optimize`, and dropping frames `gif chop drop every <N>`.

### roll

![before](doc/eggplant.png)![after](doc/roll.gif)

```text
${USAGE_roll}
```

### wobble

![before](doc/eggplant.png)![after](doc/wobble.gif)

```text
${USAGE_wobble}
```

### pulse

![before](doc/eggplant.png)![after](doc/pulse.gif)

```text
${USAGE_pulse}
```

### zoom

![before](doc/eggplant.png)![after](doc/zoom.gif)

```text
${USAGE_zoom}
```

### shake

![before](doc/eggplant.png)![after](doc/shake.gif)

```text
${USAGE_shake}
```

### woke

![before](doc/yeet.png)![after](doc/woke.gif)

```text
${USAGE_woke}
```

### fried

![before](doc/yeet.png)![after](doc/fried.gif)

```text
${USAGE_fried}
```

### hue

![before](doc/eggplant.png)![after](doc/hue.gif)

```text
${USAGE_hue}
```

### tint

![before](doc/eggplant.png)![after](doc/tint.gif)

```text
${USAGE_tint}
```

### resize

```text
${USAGE_resize}
```

### crop

```text
${USAGE_crop}
```

### optimize

```text
${USAGE_optimize}
```

### compose

![before](doc/yeet.png)![before](doc/eggplant.png)![after](doc/compose.gif)

```text
${USAGE_compose}
```

### crowd

![before](doc/wobble.gif)![after](doc/crowd.gif)

```text
${USAGE_crowd}
```

### erase

![before](doc/skeledance.gif)![after](doc/erase.gif)

```text
${USAGE_erase}
```

### chop

```text
${USAGE_chop}
```

### text

![before](doc/gunther.jpg)![after](doc/gunther.gif)
> woke | text | fried

```text
${USAGE_text}
```

### emoji

![example](doc/emoji-terminal.gif)
> emoji | compose <(emoji) | compose <(emoji) | wobble | fried

```text
${USAGE_emoji}
```

### rain

![example](doc/rain.gif)

> emoji | rain

![example](doc/rain-thonk.gif)

> emoji | roll | rain <(emoji) <(emoji)

![example](doc/rain-scream.gif)

> emoji | pulse | rain <(emoji) | compose | fried

```text
${USAGE_rain}
```

### cat

```text
${USAGE_cat}
```

### meta


![input](doc/yeet.gif)
```sh
$ <doc/yeet.gif gif meta show

[2018-10-05T13:08:57+02:00] gif fried
[2018-10-05T13:08:58+02:00] gif wobble
[2018-10-05T13:08:58+02:00] gif crop
[2018-10-05T13:08:58+02:00] gif optimize -x 0

$ <doc/yeet.gif gif meta show -p

gif fried | gif wobble | gif crop | gif optimize -x 0

$ <doc/yeet.gif gif meta show --raw

{"appName":"gif","timestamp":"2018-10-05T13:08:57+02:00","args":["fried"],"version":"1.0.0-244bcd73467a0979cb872f0e90ba8a69d4764410"}
{"appName":"gif","timestamp":"2018-10-05T13:08:58+02:00","args":["wobble"],"version":"1.0.0-244bcd73467a0979cb872f0e90ba8a69d4764410"}
{"appName":"gif","timestamp":"2018-10-05T13:08:58+02:00","args":["crop"],"version":"1.0.0-244bcd73467a0979cb872f0e90ba8a69d4764410"}
{"appName":"gif","timestamp":"2018-10-05T13:08:58+02:00","args":["optimize","-x","0"],"version":"1.0.0-244bcd73467a0979cb872f0e90ba8a69d4764410"}
```

```text
${USAGE_meta}
```

## Licensing

- `yeetgif` itself & any original code: [MIT License](LICENSE)
- [Modified copy](pkg/imaging) of `github.com/disintegration/imaging`: [MIT License](pkg/imaging/LICENSE)
- [Modified copy](pkg/ggtext) of `github.com/fogleman/gg`: [MIT License](pkg/ggtext/LICENSE.md)
- [Roboto Regular TrueType Font](pkg/gifstatic/roboto.go): [Apache License 2.0](pkg/gifstatic/roboto.go-LICENSE)
- [Twemoji](pkg/gifstatic/emoji_twitter.go) by [Twitter](https://twemoji.twitter.com): [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/)
- [Modified portion](pkg/box2d) of `github.com/ByteArena/box2d` by [ByteArena (Go port)](github.com/ByteArena/box2d)/[Erin Catto (C++ original)](http://www.box2d.org): [MIT License](pkg/box2d/LICENSE.md)
