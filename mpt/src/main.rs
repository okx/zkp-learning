// extern crate trie_db;
// extern crate trie_db_test;

//use hash_db::Hasher;
//use trie_db::DBValue;
//#[warn(unused_imports)]
use reference_trie::{ExtensionLayout};
use keccak_hasher::KeccakHasher;
use memory_db::*;
use trie_db::{Trie, TrieMut, TrieDBBuilder, TrieDBMutBuilder, TrieDBIterator};




fn main() {
    let mut memdb = MemoryDB::<KeccakHasher, HashKey<_>, _>::default();
    let mut root = Default::default();
    TrieDBMutBuilder::<ExtensionLayout>::new(&mut memdb, &mut root).build().insert(b"foo", b"bar").unwrap();
    let t = TrieDBBuilder::<ExtensionLayout>::new(&mut memdb, &mut root).build();
    assert!(t.contains(b"foo").unwrap());
    assert_eq!(t.get(b"foo").unwrap().unwrap(), b"bar".to_vec());
    println!("{:?}", t.get(b"foo").unwrap());
    println!("{:?}", t.get_hash(b"foo").unwrap());
    assert_eq!(t.get(b"foo").unwrap().unwrap(), b"barr".to_vec());
    
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_root() {
        let mut memdb = MemoryDB::<KeccakHasher, HashKey<_>, _>::default();
        let mut root = Default::default();
        let mut t = TrieDBMutBuilder::<ExtensionLayout>::new(&mut memdb, &mut root).build();
        let root1 = t.root().clone();

        t.insert(b"foo", b"bar").unwrap();
        let root2 = t.root().clone();
        assert_ne!(root1, root2.clone());

        t.remove(b"foo").unwrap();
        let root3 = t.root().clone();
        assert_eq!(root1, root3);
    }

    #[test]
    fn test_get() {
        let mut memdb = MemoryDB::<KeccakHasher, HashKey<_>, _>::default();
        let mut root = Default::default();
        let mut t = TrieDBMutBuilder::<ExtensionLayout>::new(&mut memdb, &mut root).build();
        t.insert(b"foo", b"bar").unwrap();
        t.commit();
        assert!(!t.is_empty());
        assert!(t.contains(b"foo").unwrap());
        assert_eq!(t.get(b"foo").unwrap().unwrap(), b"bar".to_vec());

        t.remove(b"foo").unwrap();
        t.commit();
        assert!(t.is_empty());
        assert!(!t.contains(b"foo").unwrap());
        assert_eq!(t.get(b"foo").unwrap().unwrap_or_default(), b"".to_vec());
    }

    #[test]
    fn test_iterator() {
        let mut memdb = MemoryDB::<KeccakHasher, HashKey<_>, _>::default();
        let mut root:[u8; 32] = Default::default();
        let mut t = TrieDBMutBuilder::<ExtensionLayout>::new(&mut memdb, &mut root).build();
        t.insert(b"acc0", b"Alice").unwrap();
        t.insert(b"acc1", b"Bob").unwrap();
        t.insert(b"acc2", b"Carol").unwrap();
        t.insert(b"acc3", b"David").unwrap();
        t.commit();

        drop(t);

        let t = TrieDBBuilder::<ExtensionLayout>::new(&memdb, &root).build();
        assert!(t.contains(b"acc0").unwrap()); 

        let iter = TrieDBIterator::new_prefixed(&t, b"acc").unwrap();
        for x in iter {
            let (key, value) = x.unwrap();
            println!("{:?}, {:?}", key, value);
        }
    }
}
