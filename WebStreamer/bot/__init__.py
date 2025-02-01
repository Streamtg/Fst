from pyrogram import Client  # ✅ Importación correcta
from ..vars import Var  # Asegura que 'Var' se importe correctamente
StreamBot = Client(
    "WebStreamer",  # ✅ Nombre de sesión sin argumento explícito
    api_id=Var.API_ID,
    api_hash=Var.API_HASH,
    bot_token=Var.BOT_TOKEN,
    sleep_threshold=Var.SLEEP_THRESHOLD,
    workers=Var.WORKERS
)
