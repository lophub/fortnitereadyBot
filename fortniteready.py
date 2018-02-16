import discord
import discordtoken
from discord.ext import commands
from discordtoken import token
import threading

client = discord.Client()


@client.event
async def on_ready():
    print('Logged in as')
    print(client.user.name)
    print(client.user.id)
    print('------')

@client.event
async def on_message(message):
    if message.content.startswith('!test'):
        counter = 0
        tmp = await client.send_message(message.channel, 'Calculating messages...')
        async for log in client.logs_from(message.channel, limit=100):
            if log.author == message.author:
                counter += 1

        await client.edit_message(tmp, 'You have {} messages.'.format(counter))
    elif message.content.startswith('!sleep'):
        await asyncio.sleep(5)
        await client.send_message(message.channel, 'Done sleeping')


def sideThread():
    # Do what you want here
    print('Worker')
    return
 
t = threading.Thread(target=sideThread)
t.start()


client.run(token)




# def botthread():
#     """thread worker function"""
#     print ('Worker')
#     bot.run(token)
#     return
