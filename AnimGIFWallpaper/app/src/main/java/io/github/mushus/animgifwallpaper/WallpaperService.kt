package io.github.mushus.animgifwallpaper

import android.graphics.Color
import android.graphics.Paint
import android.service.wallpaper.WallpaperService
import android.graphics.Movie

class WallpaperService: WallpaperService() {
    override fun onCreateEngine(): Engine {
        return WallpaperEngine()
    }

    inner class WallpaperEngine: WallpaperService.Engine() {

        private val animation: AnimationManager?

        init {
            val raw = resources.openRawResource(R.raw.testimage)
            val gif = Movie.decodeStream(raw)
            animation = AnimationManager(gif)
        }

        override fun onVisibilityChanged(visible: Boolean) {
            super.onVisibilityChanged(visible)
            if (visible) {
                drawFrame()
            }
        }

        fun drawFrame() {


            val canvas = surfaceHolder.lockCanvas()
            canvas?.let { canvas ->
                val paint = Paint()
                paint.color = Color.BLACK
                canvas.drawText("Hello Custom View!", 50f, 50f, paint)
                surfaceHolder.unlockCanvasAndPost(canvas)
            }
        }
    }
}