pub fn production_rate_per_hour(speed: u8) -> f64 {
    
    let base_rate: f64 = 221.0;
    let new_speed = speed as f64;
    let success_rate = match speed {
        1..=4 => 1.0,
        5..=8 => 0.9,
        _ => 0.77,
    };
    return new_speed * success_rate * base_rate;
}

pub fn working_items_per_minute(speed: u8) -> u32 {

    return (production_rate_per_hour(speed)/60.0).floor() as u32;
   
}
