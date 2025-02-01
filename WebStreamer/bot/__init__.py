StreamBot = Client(
    "WebStreamer",  # ✅ Nombre de sesión sin argumento explícito
    api_id=Var.API_ID,
    api_hash=Var.API_HASH,
    bot_token=Var.BOT_TOKEN,
    sleep_threshold=Var.SLEEP_THRESHOLD,
    workers=Var.WORKERS
)
