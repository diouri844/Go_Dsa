<h1 id="usage">Usage :</h1>
<h3 id="create-node">✔️ Create node:</h3>
<pre><code class="language-go">    node1 := &amp;Node{data: &quot;N1&quot;}
    node2 := &amp;Node{data: &quot;N2&quot;}
    node2 := &amp;Node{data: &quot;N3&quot;}
</code></pre>
<h3 id="create-linkedlist">✔️ Create LinkedList :</h3>
<pre><code class="language-go">    myList := LinkedList{}
</code></pre>
<h3 id="check-if-list-is-empty">✔️ Check if list is empty :</h3>
<pre><code class="language-go">    myList.isEmpty() // return true
</code></pre>
<h3 id="insert-node-as-header-index-0">✔️ Insert node  as header (index  0):</h3>
<pre><code class="language-go">    myList.inserHeadNode(node1)
</code></pre>
<h3 id="insert-node-after-index">✔️ Insert node after index :</h3>
<pre><code class="language-go">    myList.insertAfter(0, node3)
</code></pre>
<h3 id="insert-node-befor-index">✔️ Insert node befor index :</h3>
<pre><code class="language-go">    myList.insertBefor(2, node2)
</code></pre>
<h3 id="check-if-node-in-list">✔️ Check if node in list :</h3>
<pre><code class="language-go">    myList.heckNode(node2) // return true
</code></pre>
<h3 id="print-out-all-the-nodes">✔️ Print out all the nodes :</h3>
<pre><code class="language-go">    myList.showNodes()
</code></pre>
<h3 id="get-node-by-index">✔️ Get node by index :</h3>
<pre><code class="language-go">    myList.getNode(1)
</code></pre>
<h3 id="delete-the-head-node">✔️ Delete the head node :</h3>
<pre><code class="language-go">    myList.deleteHader()
</code></pre>
<h3 id="delete-node-by-index">✔️ Delete node by index  :</h3>
<pre><code class="language-go">    myList.deleteNode(index)
</code></pre>
<h3 id="delete-node-after-node">✔️ Delete node after node-target :</h3>
<pre><code class="language-go">    myList.deleteAfter(node)
</code></pre>
<h3 id="convert-to-array">✔️ Convert to array : </h3>
<pre><code class="language-go"> array := myList.toArray()
</code></pre>