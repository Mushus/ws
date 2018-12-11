package io.github.mushus.animgifwallpaper

import android.view.Choreographer
import android.graphics.Movie

class AnimationManager(movie: Movie): Choreographer.FrameCallback {

    private val choreographer: Choreographer
    private val movie: Movie = movie

    init {
        choreographer = Choreographer.getInstance();
    }

    fun start() {
        choreographer.postFrameCallback(this)
    }

    fun stop() {
        choreographer.removeFrameCallback(this)
    }

    public override fun doFrame(var1: Long) {
        choreographer.postFrameCallback(this)
    }
}