import discord
import discordtoken
from discord.ext import commands
from discordtoken import token
import threading
import time
import readline
import asyncio


test_channel = "testing-ground"

loop = asyncio.new_event_loop()
bot = discord.Client()
bot_token = True
message_queue = asyncio.Queue()
# Testing ground
channel_id = "413813088018497536"

def bot_thread(loop, bot, bot_token, message_queue, channel_id):
    asyncio.set_event_loop(loop)

    @bot.event
    async def on_ready():
        while True:
            data = await message_queue.get()
            event = data[0]
            message = data[1]
            channel_id = data[2]

            try:
                print(message)
                await bot.send_message(bot.get_channel(channel_id), message)
            except:
                pass

            event.set()

    bot.run(token, bot = bot_token)

thread = threading.Thread(target = bot_thread, args = (loop, bot, bot_token, message_queue, channel_id), daemon = True)
thread.start()

def send(channel_id, message):
    event = threading.Event()

    message_queue.put_nowait([event, message, channel_id])

    event.wait()

print('Bot logging in...')

try:
    kek = False
    while not kek:
        time.sleep(0.1)

        if bot._is_ready.is_set(): # wait until the ready event
            while True:
                try:
                    message = input('Message to send > ')
                except:
                    break
                else:
                    send(channel_id, message)

            kek = True
except KeyboardInterrupt:
    pass

# # http://discordpy.readthedocs.io/en/latest/api.html#discord.Client.get_all_members
# client = discord.Client()


# @client.event
# async def on_ready():
#     print('Logged in as')
#     print(client.user.name)
#     print(client.user.id)
#     print('------')
#     t.start()

# @client.event
# async def on_message(message):
#     if message.content.startswith('!test'):
#         counter = 0
#         print(message.channel)
#         tmp = await client.send_message(message.channel, 'Calculating messages...')
#         async for log in client.logs_from(message.channel, limit=100):
#             if log.author == message.author:
#                 counter += 1

#         await client.edit_message(tmp, 'You have {} messages.'.format(counter))
#     elif message.content.startswith('!sleep'):
#         await asyncio.sleep(5)
#         await client.send_message(message.channel, 'Done sleeping')


# async def send_to_test_channel(message):
#     print(message)
#     await client.send_message(test_channel, message)

# def sideThread():
#     # Do what you want here
#     send_to_test_channel("Hey!")
#     print('Worker')
#     return
 
# t = threading.Thread(target=sideThread)


# await client.login(token)
# await client.connect()
# # client.start(token)

# while True:
#     print("...")
#     time.sleep(5)




# # def botthread():
# #     """thread worker function"""
# #     print ('Worker')
# #     bot.run(token)
# #     return
