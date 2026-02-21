<script>
  import { fetchWithPDP } from '../../../libs/js/index.js';

  let messages = $state([
    { role: 'ai', content: 'Hello! I am a PDP-compliant assistant. Adjust the privacy level below to see how the protocol works.' }
  ]);
  let input = $state('');
  let pdpLevel = $state(0); 
  let isSending = $state(false);

  async function sendMessage(e) {
    if (e) e.preventDefault();
    if (!input.trim() || isSending) return;
    
    const userMessage = { role: 'user', content: input };
    messages = [...messages, userMessage];
    const currentInput = input;
    input = '';
    isSending = true;

    try {
      await fetchWithPDP('/api/chat', {
        method: 'POST',
        body: JSON.stringify({ messages: [...messages] })
      }, pdpLevel);

      messages = [...messages, { role: 'ai', content: 'Response received. Check the right panel for the PDP header logs.' }];
    } catch (e) {
      console.error("PDP Request Failed", e);
    } finally {
      isSending = false;
    }
  }
</script>

<main class="flex h-screen bg-gray-50 font-sans text-gray-900">
  <section class="flex flex-col flex-1 border-r border-gray-200">
    <header class="p-4 bg-white border-b border-gray-200 flex justify-between items-center">
      <h1 class="font-bold text-lg tracking-tight">PDP Messenger</h1>
      <div class="flex gap-2 bg-gray-100 p-1 rounded-lg">
        {#each [0, 1, 2] as level}
          <button 
            type="button"
            onclick={() => pdpLevel = level}
            class="px-3 py-1 rounded-md text-sm transition-all {pdpLevel === level ? 'bg-white shadow-sm font-bold' : 'text-gray-500 hover:text-gray-700'}"
          >
            Level {level}
          </button>
        {/each}
      </div>
    </header>

    <div class="flex-1 overflow-y-auto p-6 space-y-4">
      {#each messages as msg}
        <div class="flex {msg.role === 'user' ? 'justify-end' : 'justify-start'}">
          <div class="max-w-md p-3 rounded-2xl {msg.role === 'user' ? 'bg-black text-white' : 'bg-white border border-gray-200 shadow-sm'}">
            {msg.content}
          </div>
        </div>
      {/each}
    </div>

    <form onsubmit={sendMessage} class="p-4 bg-white border-t border-gray-200 flex gap-2">
      <input 
        bind:value={input} 
        placeholder="Type a prompt..." 
        class="flex-1 bg-gray-100 border-none rounded-xl px-4 focus:ring-2 focus:ring-black transition-all"
      />
      <button type="submit" class="bg-black text-white px-6 py-2 rounded-xl font-medium hover:opacity-80 transition-opacity">
        Send
      </button>
    </form>
  </section>

  <aside class="w-80 bg-gray-900 text-gray-300 p-6 overflow-hidden flex flex-col">
    <h2 class="text-xs font-semibold uppercase tracking-widest text-gray-500 mb-4">Protocol Inspector</h2>
    
    <div class="space-y-6">
      <div>
        <span class="text-[10px] text-gray-500 uppercase block">Active Header</span>
        <div class="mt-1 font-mono text-sm bg-black/50 p-3 rounded-lg border border-gray-800">
          <span class="text-blue-400">X-PDP-Level:</span> 
          <span class="text-green-400">{pdpLevel}</span>
        </div>
      </div>

      <div>
        <span class="text-[10px] text-gray-500 uppercase block">Level Description</span>
        <div class="mt-2 text-sm leading-relaxed">
          {#if pdpLevel === 0}
            <strong class="text-white block mb-1">Strictly Private</strong>
            The provider is prohibited from storing or training on this prompt.
          {:else if pdpLevel === 1}
            <strong class="text-white block mb-1">Local Refinement</strong>
            Data may be used for your personal fine-tuning only.
          {:else}
            <strong class="text-white block mb-1">Global Training</strong>
            Full consent granted for base model improvement.
          {/if}
        </div>
      </div>
    </div>

    <footer class="mt-auto pt-6 border-t border-gray-800">
      <p class="text-[10px] text-gray-600">Prompt Data Protocol v1.0.0</p>
    </footer>
  </aside>
</main>
