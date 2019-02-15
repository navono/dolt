# 1/17/19 Version 0.3.0
  * Merge Support
    * dolt merge &lt;branch&gt;
    * dolt merge --abort
  * Conflict display
    * dolt conflicts cat [&lt;commit&gt;] &lt;table&gt;...
  * Resolve conflicts
    * dolt conflicts resolve &lt;table&gt; &lt;key&gt;...
    * dolt conflicts resolve --ours|--theirs &lt;table&gt;...


# 1/8/19 Version 0.2.2
  * Windows support.
  * Move data to .dolt/noms directory
    * To convert existing data repositories you will need to run:
      * mkdir .dolt/noms
      * mv * .dolt/noms/
      * rm * (If you have been using this directory to store other files you may need to be more deliberate about what you delete.)
  * Bug fixes.
  * Error messaging improvements.

# 1/6/19 Version 0.2.1
  * dolt table schema command
  * --all option for dolt add
  * table headers printed when using commands cat and diff
  * Bug fixes
  * Documentation Updates

# 1/4/19 Version 0.2.0
Initial internal release of dolt for team to be able to start playing with it.  Initial documentetation:
