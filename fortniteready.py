import discord
import discordtoken
from discord.ext import commands
from discordtoken import token


bot_prefix = "!"
bot = commands.Bot(command_prefix=bot_prefix)

@bot.event
async def on_ready():
    print("Bot Online!")

@bot.command(pass_context=True)
async def ping():
    await bot.say("Pong!")


@bot.command(pass_context=True)
async def stats():
    await bot.say("Lop's better than you.")

@bot.command(pass_context=True)
async def ready():
    await bot.say("I'm ready to play!")

@bot.command(pass_context=True)
async def kagy():
    await bot.say("Gay.")

@bot.command(pass_context=True)
async def lop():
    await bot.say("Beast.")

@bot.command(pass_context=True)
async def emyrk():
    await bot.say("Gay.")

@bot.command(pass_context=True)
async def robdog364():
    await bot.say("Gay.")

@bot.command(pass_context=True)
async def guywhodoesthatthing():
    await bot.say("Gay.")

@bot.command(pass_context=True)
async def prelor():
    await bot.say("Gay.")

@bot.command(pass_context=True)
async def jackal():
    await bot.say("Gay.")

@bot.command(pass_context=True)
async def gay():
    await bot.say("Straight.")

 

bot.run(token)