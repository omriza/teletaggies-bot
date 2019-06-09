# TeleTagies Bot

TeleTagies is a telegram bot for storing and tagging links, media and other content for later use.

## What

### Desired Features

- Content Storage
  - Tagging
    - Automatic Tagging Suggestions (Heuristic Based)
      - Type
      - Domain Based
        - Domain Name
        - Extract more data from links by following them 
    - Tagging Retrospectively
  
- Querying Stored Content
  - By Tag
  - By Date
  - By Type
  - Free Text

- Get More - Machine Learning

## How

### Access

- Who can use the bot?
- Privacy

### Storage

As messages are saved in Telegram itself, I can take advantage of Telegram's free hosting to keep large volume content such as photos, videos and other files and save only metadata in the database. This has some omplications tho-

- What happens if the user deletes the chat or edit the content?
- Are Telegram messages kept forever?
- As Telegram chats are E2E encrypted, what if I change the bot user, thus losing the ability to access the messages?

### Message Metadata

#### Generic Metadata for all Content Types

- Telegram ID
- Sender Info
- Unixtime
- Chat
- OriginalSender
- OriginalChat
- OriginalUnixtime
- ReplyTo
- LastEdit

#### Specific Type Metadata

**Text Messages**
- Text

**Photos**
- Photo
  - Width
  - Height

**Videos**
- Video
  - Width
  - Height
  - Duration

**<u>Might be implemented in the future:</u>**
**Audio Messages**

- Audio / Voice
  - Duration

**Generic File Messages**
- Document
  - FileName
