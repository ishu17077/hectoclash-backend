<script>
    let isLogin = true;
</script>

<div class="h-fit flex items-center justify-center bg-black text-green-500">
    <div class="w-full max-w-md bg-black border-4 border-green-500 rounded-lg p-6 space-y-8">
        <!-- Toggle Buttons -->
        <div class="flex justify-center space-x-4">
            <button 
                class="px-4 py-2 text-sm font-bold retro-font border-2 rounded-md 
                       focus:outline-wh= hover:text-black 
                       {isLogin ? 'bg-green-500 text-black' : 'bg-black text-green-500'}"
                onclick={() => isLogin = true}
            >
                LOG IN
            </button>
            <button 
                class="px-4 py-2 text-sm font-bold retro-font border-2 rounded-md 
                       focus:outline-none focus:ring-2 focus:ring-green-500 
                       hover:bg-green-500 hover:text-black 
                       {!isLogin ? 'bg-green-500 text-black' : 'bg-black text-green-500'}"
                onclick={() => isLogin = false}
            >
                SIGN UP
            </button>
        </div>

        {#if isLogin}
        <form method="POST" action="?/login" class="space-y-4">
            <h2 class="text-3xl font-bold text-center retro-font">LOG IN</h2>
            <div>
                <label class="block text-sm font-bold retro-font">EMAIL</label>
                <input 
                    name="email" 
                    type="email" 
                    required 
                    class="mt-1 block w-full px-3 py-2 border-2 border-green-500 bg-black text-green-500 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 retro-font"
                >
            </div>
            <div>
                <label class="block text-sm font-bold retro-font">PASSWORD</label>
                <input 
                    name="password" 
                    type="password" 
                    required 
                    class="mt-1 block w-full px-3 py-2 border-2 border-green-500 bg-black text-green-500 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 retro-font"
                >
            </div>
            <button 
                type="submit" 
                class="w-full bg-green-500 text-black py-2 px-4 rounded-md hover:bg-green-400 focus:outline-none focus:ring-2 focus:ring-green-300 retro-font"
            >
                LOG IN
            </button>
        </form>
        {/if}

        <!-- Signup Form -->
        {#if !isLogin}
        <form method="POST" action="?/signin" class="space-y-4">
            <h2 class="text-3xl font-bold text-center retro-font">SIGN UP</h2>
            <div>
                <label class="block text-sm font-bold retro-font">NAME</label>
                <input 
                    name="username" 
                    type="text" 
                    required 
                    class="mt-1 block w-full px-3 py-2 border-2 border-green-500 bg-black text-green-500 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 retro-font"
                >
            </div>
            <div>
                <label class="block text-sm font-bold retro-font">EMAIL</label>
                <input 
                    name="email" 
                    type="email" 
                    required 
                    class="mt-1 block w-full px-3 py-2 border-2 border-green-500 bg-black text-green-500 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 retro-font"
                >
            </div>
            <div>
                <label class="block text-sm font-bold retro-font">PASSWORD</label>
                <input 
                    name="password" 
                    type="password" 
                    required 
                    class="mt-1 block w-full px-3 py-2 border-2 border-green-500 bg-black text-green-500 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 retro-font"
                >
            </div>
            <button 
                type="submit" 
                class="w-full bg-green-500 text-black py-2 px-4 rounded-md hover:bg-green-400 focus:outline-none focus:ring-2 focus:ring-green-300 retro-font"
            >
                SIGN UP
            </button>
        </form>
        {/if}
    </div>
</div>