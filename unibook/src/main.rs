use alloy::{
    primitives::{address, U256}, providers::{Provider, ProviderBuilder}
};

use eyre::Result;
use std::{
    error::Error
}

const RPC_URL: &str = std::env::var("RPC_URL");

async fn fetch_prices() -> Result<(), Box<dyn Error>> {
    let provider = ProviderBuilder::new().on_http(RPC_URL).await?;
    let storage_slot = U256::from(0);
    let pool_address = address!("0xd0b53D9277642d899DF5C87A3966A349A798F224");
    let storage = provider.get_storage_at(pool_address, storage_slot).await?;

    println!("Slot 0: {storage:?}");

    Ok(())
}

#[tokio::main]
fn main() {
    fetch_prices();
}
